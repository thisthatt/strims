import { Error } from "@memelabs/protobuf/lib/apis/strims/rpc/rpc";
import { Base64 } from "js-base64";
import React, { useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { MdChevronLeft, MdChevronRight } from "react-icons/md";
import { Link, Redirect, Switch, useHistory, useParams } from "react-router-dom";
import Select from "react-select";
import { useAsync } from "react-use";

import {
  Emote,
  EmoteFileType,
  EmoteImage,
  EmoteScale,
  Server,
} from "../../apis/strims/chat/v1/chat";
import {
  ImageInput,
  ImageValue,
  InputError,
  InputLabel,
  TextAreaInput,
  TextInput,
  ToggleInput,
} from "../../components/Form";
import { PrivateRoute } from "../../components/PrivateRoute";
import { useCall, useClient, useLazyCall } from "../../contexts/FrontendApi";
import { rootCertificate } from "../../lib/certificate";
import jsonutil from "../../lib/jsonutil";

interface BackLinkProps {
  to: string;
  title: string;
  description: string;
}

const BackLink: React.FC<BackLinkProps> = ({ to, title, description }) => (
  <Link className="input_label input_button" to={to}>
    <MdChevronLeft size="28" />
    <div className="input_label__body">
      <div>{title}</div>
      <div>{description}</div>
    </div>
  </Link>
);

interface ChatServerFormData {
  name: string;
  networkKey: {
    value: string;
    label: string;
  };
}

interface ChatServerFormProps {
  onSubmit: (data: ChatServerFormData) => void;
  error: Error;
  loading: boolean;
  config?: Server;
  indexLinkVisible?: boolean;
}

const ChatServerForm: React.FC<ChatServerFormProps> = ({
  onSubmit,
  error,
  loading,
  config,
  indexLinkVisible,
}) => {
  const client = useClient();

  const { handleSubmit, control, reset } = useForm<ChatServerFormData>({
    mode: "onBlur",
  });

  const { value: networkOptions } = useAsync(async () => {
    const res = await client.network.list();
    return res.networks.map((n) => ({
      value: Base64.fromUint8Array(rootCertificate(n.certificate).key),
      label: n.name,
    }));
  });

  const setValues = (config: Server) => {
    const networkKey = Base64.fromUint8Array(config.networkKey);

    reset(
      {
        name: config.room.name,
        networkKey: networkOptions.find(({ value }) => value === networkKey),
      },
      {
        keepDirty: false,
        keepIsValid: false,
      }
    );
  };

  React.useEffect(() => {
    if (networkOptions && config) {
      setValues(config);
    }
  }, [networkOptions, config]);

  return (
    <form className="thing_form" onSubmit={handleSubmit(onSubmit)}>
      {error && <InputError error={error.message || "Error creating chat server"} />}
      {indexLinkVisible && (
        <BackLink
          to="/settings/chat-servers"
          title="Servers"
          description="Some description of servers..."
        />
      )}
      <TextInput
        control={control}
        rules={{
          required: {
            value: true,
            message: "Name is required",
          },
        }}
        name="name"
        label="Name"
        placeholder="Enter a chat room name"
      />
      <InputLabel required={true} text="Network">
        <Controller
          name="networkKey"
          control={control}
          rules={{
            required: {
              value: true,
              message: "Network is required",
            },
          }}
          render={({ field, fieldState: { error } }) => (
            <>
              <Select
                {...field}
                className="input_select"
                classNamePrefix="react_select"
                placeholder="Select network"
                options={networkOptions}
              />
              <InputError error={error} />
            </>
          )}
        />
      </InputLabel>
      {config && (
        <Link
          className="input_label input_button"
          to={`/settings/chat-servers/${config.id}/emotes`}
        >
          <div className="input_label__body">
            <div>Emotes</div>
            <div>Some description of emotes...</div>
          </div>
          <MdChevronRight size="28" />
        </Link>
      )}
      <label className="input_label">
        <div className="input_label__body">
          <button className="input input_button" disabled={loading}>
            {config ? "Update Server" : "Create Server"}
          </button>
        </div>
      </label>
    </form>
  );
};

interface ChatServerTableProps {
  servers: Server[];
  onDelete: () => void;
}

const ChatServerTable: React.FC<ChatServerTableProps> = ({ servers, onDelete }) => {
  const [{ error }, deleteChatServer] = useLazyCall("chat", "deleteServer", {
    onComplete: onDelete,
  });

  if (!servers) {
    return null;
  }

  const rows = servers.map((server, i) => {
    const handleDelete = () => deleteChatServer({ id: server.id });

    return (
      <div className="thing_list__item" key={server.id.toString()}>
        <Link to={`/settings/chat-servers/${server.id}`}>{server.room.name || "no title"}</Link>
        <button className="input input_button" onClick={handleDelete}>
          delete
        </button>
        <pre>{jsonutil.stringify(server)}</pre>
      </div>
    );
  });
  return <div className="thing_list">{rows}</div>;
};

const ChatServerIndexPage: React.FC = () => {
  const [{ loading, value }, getServers] = useCall("chat", "listServers");

  if (loading) {
    return null;
  }
  if (!value?.servers.length) {
    return <Redirect to="/settings/chat-servers/new" />;
  }
  return (
    <>
      <ChatServerTable servers={value.servers} onDelete={() => getServers()} />
      <Link to="/settings/chat-servers/new">Create Server</Link>
    </>
  );
};

const ChatServerCreateFormPage: React.FC = () => {
  const [{ value }] = useCall("chat", "listServers");
  const history = useHistory();
  const [{ error, loading }, createChatServer] = useLazyCall("chat", "createServer", {
    onComplete: (res) => history.replace(`/settings/chat-servers/${res.server.id}`),
  });

  const onSubmit = (data: ChatServerFormData) =>
    createChatServer({
      networkKey: Base64.toUint8Array(data.networkKey.value),
      room: {
        name: data.name,
      },
    });

  return (
    <ChatServerForm
      onSubmit={onSubmit}
      error={error}
      loading={loading}
      indexLinkVisible={!!value?.servers.length}
    />
  );
};

const ChatServerEditFormPage: React.FC = () => {
  const { serverId } = useParams<{ serverId: string }>();
  const [getRes] = useCall("chat", "getServer", { args: [{ id: BigInt(serverId) }] });

  const [updateRes, updateChatServer] = useLazyCall("chat", "updateServer");

  const onSubmit = (data: ChatServerFormData) =>
    updateChatServer({
      id: BigInt(serverId),
      networkKey: Base64.toUint8Array(data.networkKey.value),
      room: {
        name: data.name,
      },
    });

  return (
    <ChatServerForm
      onSubmit={onSubmit}
      error={getRes.error || updateRes.error}
      loading={getRes.loading || updateRes.loading}
      config={getRes.value?.server}
      indexLinkVisible={true}
    />
  );
};

interface ChatEmoteFormData {
  name: string;
  image: ImageValue;
  scale: {
    value: EmoteScale;
    label: string;
  };
  css: string;
  animated: boolean;
  animationFrameCount: string;
  animationDuration: string;
  animationIterationCount: string;
}

interface ChatEmoteFormProps {
  onSubmit: (data: ChatEmoteFormData) => void;
  error: Error;
  loading: boolean;
  serverId: bigint;
  values?: ChatEmoteFormData;
  indexLinkVisible: boolean;
}

const ChatEmoteForm: React.FC<ChatEmoteFormProps> = ({
  onSubmit,
  error,
  loading,
  serverId,
  values = {},
  indexLinkVisible,
}) => {
  const { handleSubmit, control, watch } = useForm<ChatEmoteFormData>({
    mode: "onBlur",
    defaultValues: {
      scale: {
        value: EmoteScale.EMOTE_SCALE_1X,
        label: "1x",
      },
      ...values,
    },
  });

  const animated = watch("animated");

  return (
    <form className="thing_form" onSubmit={handleSubmit(onSubmit)}>
      {error && <InputError error={error.message || "Error creating chat server"} />}
      {indexLinkVisible ? (
        <BackLink
          to={`/settings/chat-servers/${serverId}/emotes`}
          title="Emotes"
          description="Some description of emotes..."
        />
      ) : (
        <BackLink
          to={`/settings/chat-servers/${serverId}`}
          title="Server"
          description="Some description of server..."
        />
      )}
      <TextInput
        control={control}
        rules={{
          required: {
            value: true,
            message: "Name is required",
          },
        }}
        name="name"
        label="Name"
        placeholder="Enter a emote name"
      />
      <InputLabel required={true} text="Image">
        <ImageInput control={control} name="image" />
      </InputLabel>
      <InputLabel required={true} text="Scale">
        <Controller
          name="scale"
          control={control}
          render={({ field, fieldState: { error } }) => (
            <>
              <Select
                {...field}
                className="input_select"
                classNamePrefix="react_select"
                placeholder="Scale"
                options={[
                  {
                    value: EmoteScale.EMOTE_SCALE_1X,
                    label: "1x",
                  },
                  {
                    value: EmoteScale.EMOTE_SCALE_2X,
                    label: "2x",
                  },
                  {
                    value: EmoteScale.EMOTE_SCALE_4X,
                    label: "4x",
                  },
                ]}
              />
              <InputError error={error} />
            </>
          )}
        />
      </InputLabel>
      <TextAreaInput control={control} label="css" name="css" />
      <ToggleInput control={control} label="animated" name="animated" />
      <TextInput
        control={control}
        label="frame count"
        name="animationFrameCount"
        disabled={!animated}
      />
      <TextInput control={control} label="duration" name="animationDuration" disabled={!animated} />
      <TextInput
        control={control}
        label="loops"
        name="animationIterationCount"
        disabled={!animated}
      />
      <label className="input_label">
        <div className="input_label__body">
          <button className="input input_button" disabled={loading}>
            {values ? "Update Emote" : "Create Emote"}
          </button>
        </div>
      </label>
    </form>
  );
};

interface ImageProps {
  src: EmoteImage;
}

const Image: React.FC<ImageProps> = ({ src }) => {
  const [url] = useState(() =>
    URL.createObjectURL(new Blob([src.data], { type: fileTypeToMimeType(src.fileType) }))
  );
  useEffect(() => () => URL.revokeObjectURL(url));

  return <img srcSet={`${url} ${scaleToDOMScale(src.scale)}`} />;
};

interface ChatEmoteTableProps {
  serverId: bigint;
  emotes: Emote[];
  onDelete: () => void;
}

const ChatEmoteTable: React.FC<ChatEmoteTableProps> = ({ serverId, emotes, onDelete }) => {
  const [, deleteChatEmote] = useLazyCall("chat", "deleteEmote", { onComplete: onDelete });

  if (!emotes) {
    return null;
  }

  const rows = emotes.map((emote, i) => {
    const handleDelete = () => deleteChatEmote({ serverId, id: emote.id });

    return (
      <div className="thing_list__item" key={emote.id.toString()}>
        <Image src={emote.images[0]} />
        <Link to={`/settings/chat-servers/${serverId}/emotes/${emote.id}`}>{emote.name}</Link>
        <button className="input input_button" onClick={handleDelete}>
          delete
        </button>
      </div>
    );
  });
  return (
    <div className="thing_list">
      <BackLink
        to={`/settings/chat-servers/${serverId}`}
        title="Server"
        description="Some description of server..."
      />
      {rows}
    </div>
  );
};

const ChatServerEmotePage: React.FC = () => {
  const { serverId } = useParams<{ serverId: string }>();
  const [{ loading, value }, getEmotes] = useCall("chat", "listEmotes", {
    args: [{ serverId: BigInt(serverId) }],
  });

  if (loading) {
    return null;
  }
  if (!value?.emotes.length) {
    return <Redirect to={`/settings/chat-servers/${serverId}/emotes/new`} />;
  }
  return (
    <>
      <ChatEmoteTable
        serverId={BigInt(serverId)}
        emotes={value.emotes}
        onDelete={() => getEmotes()}
      />
      <Link to={`/settings/chat-servers/${serverId}/emotes/new`}>Create Emote</Link>
    </>
  );
};

const ChatEmoteCreateFormPage: React.FC = () => {
  const { serverId } = useParams<{ serverId: string }>();
  const [{ value }] = useCall("chat", "listEmotes", {
    args: [{ serverId: BigInt(serverId) }],
  });
  const history = useHistory();
  const [{ error, loading }, createChatEmote] = useLazyCall("chat", "createEmote", {
    onComplete: () => history.replace(`/settings/chat-servers/${serverId}/emotes`),
  });

  const onSubmit = (data: ChatEmoteFormData) =>
    createChatEmote({
      serverId: BigInt(serverId),
      name: data.name,
      images: [
        {
          data: data.image.data,
          fileType: mimeTypeToFileType(data.image.type),
          height: data.image.height,
          width: data.image.width,
          scale: data.scale.value,
        },
      ],
      css: data.css,
      animation: data.animated
        ? {
            frameCount: parseFloat(data.animationFrameCount),
            durationMs: parseFloat(data.animationDuration),
            iterationCount: parseFloat(data.animationIterationCount),
          }
        : null,
    });

  return (
    <ChatEmoteForm
      onSubmit={onSubmit}
      error={error}
      loading={loading}
      serverId={BigInt(serverId)}
      indexLinkVisible={!!value?.emotes.length}
    />
  );
};

const mimeTypeToFileType = (type: string): EmoteFileType => {
  switch (type) {
    case "image/png":
      return EmoteFileType.FILE_TYPE_PNG;
    case "image/gif":
      return EmoteFileType.FILE_TYPE_GIF;
    default:
      return EmoteFileType.FILE_TYPE_UNDEFINED;
  }
};

const fileTypeToMimeType = (type: EmoteFileType): string => {
  switch (type) {
    case EmoteFileType.FILE_TYPE_PNG:
      return "image/png";
    case EmoteFileType.FILE_TYPE_GIF:
      return "image/gif";
    case EmoteFileType.FILE_TYPE_UNDEFINED:
      return "application/octet-stream";
  }
};

const scaleToDOMScale = (type: EmoteScale): string => {
  switch (type) {
    case EmoteScale.EMOTE_SCALE_1X:
      return "1x";
    case EmoteScale.EMOTE_SCALE_2X:
      return "2x";
    case EmoteScale.EMOTE_SCALE_4X:
      return "4x";
  }
};

const ChatEmoteEditFormPage: React.FC = () => {
  const { serverId, emoteId } = useParams<{ serverId: string; emoteId: string }>();
  const [{ value, ...getRes }] = useCall("chat", "getEmote", { args: [{ id: BigInt(emoteId) }] });

  const [updateRes, updateChatEmote] = useLazyCall("chat", "updateEmote");

  const onSubmit = (data: ChatEmoteFormData) =>
    updateChatEmote({
      serverId: BigInt(serverId),
      id: BigInt(emoteId),
      name: data.name,
      images: [
        {
          data: data.image.data,
          fileType: mimeTypeToFileType(data.image.type),
          height: data.image.height,
          width: data.image.width,
          scale: data.scale.value,
        },
      ],
      css: data.css,
      animation: data.animated
        ? {
            frameCount: parseFloat(data.animationFrameCount),
            durationMs: parseFloat(data.animationDuration),
            iterationCount: parseFloat(data.animationIterationCount),
          }
        : null,
    });

  if (getRes.loading) {
    return null;
  }

  const { emote } = value;

  return (
    <ChatEmoteForm
      onSubmit={onSubmit}
      error={getRes.error || updateRes.error}
      loading={getRes.loading || updateRes.loading}
      values={{
        name: emote.name,
        image: {
          data: emote.images[0].data,
          type: fileTypeToMimeType(emote.images[0].fileType),
          height: emote.images[0].height,
          width: emote.images[0].width,
        },
        scale: {
          value: emote.images[0].scale,
          label: scaleToDOMScale(emote.images[0].scale),
        },
        css: emote.css,
        animated: !!emote.animation,
        animationFrameCount: emote.animation?.frameCount.toString(),
        animationDuration: emote.animation?.durationMs.toString(),
        animationIterationCount: emote.animation?.iterationCount.toString(),
      }}
      serverId={BigInt(serverId)}
      indexLinkVisible={true}
    />
  );
};

const ChatServersPage: React.FC = () => (
  <main className="network_page">
    <Switch>
      <PrivateRoute path="/settings/chat-servers" exact component={ChatServerIndexPage} />
      <PrivateRoute path="/settings/chat-servers/new" exact component={ChatServerCreateFormPage} />
      <PrivateRoute
        path="/settings/chat-servers/:serverId"
        exact
        component={ChatServerEditFormPage}
      />
      <PrivateRoute
        path="/settings/chat-servers/:serverId/emotes"
        exact
        component={ChatServerEmotePage}
      />
      <PrivateRoute
        path="/settings/chat-servers/:serverId/emotes/new"
        exact
        component={ChatEmoteCreateFormPage}
      />
      <PrivateRoute
        path="/settings/chat-servers/:serverId/emotes/:emoteId"
        exact
        component={ChatEmoteEditFormPage}
      />
    </Switch>
  </main>
);

export default ChatServersPage;
