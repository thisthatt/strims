// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import React from "react";
import { useTranslation } from "react-i18next";
import { Link, Navigate, useParams } from "react-router-dom";
import { useTitle } from "react-use";

import { Emote, ListEmotesRequest } from "../../../apis/strims/chat/v1/chat";
import {
  MenuCell,
  MenuItem,
  MenuLink,
  Table,
  TableMenu,
  TableTitleBar,
} from "../../../components/Settings/Table";
import { useCall, useLazyCall } from "../../../contexts/FrontendApi";

export interface ChatEmoteTableProps {
  serverId: bigint;
  emotes: Emote[];
  onDelete: () => void;
}

const ChatEmoteTable: React.FC<ChatEmoteTableProps> = ({ serverId, emotes, onDelete }) => {
  const [, deleteChatEmote] = useLazyCall("chatServer", "deleteEmote", { onComplete: onDelete });

  if (!emotes) {
    return null;
  }

  const rows = emotes.map((emote) => {
    const handleDelete = () => deleteChatEmote({ serverId, id: emote.id });

    return (
      <tr key={emote.id.toString()}>
        <td>
          <Link to={`/settings/chat-servers/${serverId}/emotes/${emote.id}`}>{emote.name}</Link>
        </td>
        <MenuCell>
          <MenuItem label="Delete" onClick={handleDelete} />
        </MenuCell>
      </tr>
    );
  });
  return (
    <Table>
      <thead>
        <tr>
          <th>Title</th>
          <th></th>
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </Table>
  );
};

const ChatEmoteList: React.FC = () => {
  const { t } = useTranslation();
  useTitle(t("settings.chat.title"));

  const { serverId } = useParams<"serverId">();
  const [{ loading, value }, getEmotes] = useCall("chatServer", "listEmotes", {
    args: [
      {
        serverId: BigInt(serverId),
        parts: [ListEmotesRequest.Part.PART_META],
      },
    ],
  });

  if (loading) {
    return null;
  }
  if (!value?.emotes.length) {
    return <Navigate to={`/settings/chat-servers/${serverId}/emotes/new`} />;
  }
  return (
    <>
      <TableTitleBar label="Emotes" backLink={`/settings/chat-servers/${serverId}`}>
        <TableMenu label="Create">
          <MenuLink label="Create Emote" to={`/settings/chat-servers/${serverId}/emotes/new`} />
        </TableMenu>
      </TableTitleBar>
      <ChatEmoteTable
        serverId={BigInt(serverId)}
        emotes={value.emotes}
        onDelete={() => getEmotes()}
      />
    </>
  );
};

export default ChatEmoteList;
