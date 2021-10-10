import { Error } from "@memelabs/protobuf/lib/apis/strims/rpc/rpc";
import React, { useEffect } from "react";
import { useForm } from "react-hook-form";

import { Network } from "../../../apis/strims/network/v1/network";
import { InputError, TextInput, ToggleInput } from "../../../components/Form";
import BackLink from "./BackLink";

export interface NetworkFormData {
  angelthumpEnable: boolean;
  twitchEnable: boolean;
  twitchClientId: string;
  twitchClientSecret: string;
  youtubeEnable: boolean;
  youtubePublicApiKey: string;
  swarmEnable: boolean;
}

export interface NetworkFormProps {
  onSubmit: (data: NetworkFormData) => void;
  error: Error;
  loading: boolean;
  values?: NetworkFormData;
  indexLinkVisible?: boolean;
}

const NetworkForm: React.FC<NetworkFormProps> = ({
  onSubmit,
  error,
  loading,
  values = {},
  indexLinkVisible,
}) => {
  const { handleSubmit, control, watch, clearErrors } = useForm<NetworkFormData>({
    mode: "onBlur",
    defaultValues: values,
  });

  const { twitchEnable, youtubeEnable } = watch();
  useEffect(() => clearErrors(["twitchClientId", "twitchClientSecret"]), [twitchEnable]);
  useEffect(() => clearErrors(["youtubePublicApiKey"]), [youtubeEnable]);

  return (
    <form className="thing_form" onSubmit={handleSubmit(onSubmit)}>
      {error && <InputError error={error.message || "Error creating tag"} />}
      {indexLinkVisible && (
        <BackLink
          to="/settings/networks"
          title="Networks"
          description="Some description of networks..."
        />
      )}
      <ToggleInput control={control} name="angelthumpEnable" label="Allow AngelThump embed" />
      <ToggleInput control={control} name="twitchEnable" label="Allow Twitch embed" />
      <TextInput
        control={control}
        rules={{
          required: {
            value: twitchEnable,
            message: "Twitch client ID",
          },
        }}
        name="twitchClientId"
        autoComplete="off"
        label="Twitch client ID"
        placeholder="Enter a Twitch client ID"
      />
      <TextInput
        control={control}
        rules={{
          required: {
            value: twitchEnable,
            message: "Twitch client secret",
          },
        }}
        name="twitchClientSecret"
        autoComplete="off"
        type="password"
        label="Twitch client Secret"
        placeholder="Enter a Twitch client secret"
      />
      <ToggleInput control={control} name="youtubeEnable" label="Allow YouTube embed" />
      <TextInput
        control={control}
        rules={{
          required: {
            value: youtubeEnable,
            message: "YouTube public API key required",
          },
        }}
        name="youtubePublicApiKey"
        autoComplete="off"
        type="password"
        label="YouTube public API key"
        placeholder="Enter a YouTube public API key"
      />
      <ToggleInput control={control} name="swarmEnable" label="Allow swarm embed" />
      <label className="input_label">
        <div className="input_label__body">
          <button className="input input_button" disabled={loading}>
            {values ? "Update Network" : "Create Network"}
          </button>
        </div>
      </label>
    </form>
  );
};

export default NetworkForm;
