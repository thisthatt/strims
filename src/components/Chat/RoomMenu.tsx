// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import "./RoomMenu.scss";

import clsx from "clsx";
import date from "date-and-time";
import { Base64 } from "js-base64";
import React, { forwardRef, useContext, useMemo, useState } from "react";
import Scrollbars from "react-custom-scrollbars-2";
import { HiOutlineUser } from "react-icons/hi";

import { WhisperThread } from "../../apis/strims/chat/v1/chat";
import * as directoryv1 from "../../apis/strims/network/v1/directory/directory";
import { Network } from "../../apis/strims/network/v1/network";
import { ThreadProviderProps, useChat } from "../../contexts/Chat";
import { DirectoryContext, DirectoryUser } from "../../contexts/Directory";
import { NetworkContext } from "../../contexts/Network";
import { useStableCallback } from "../../hooks/useStableCallback";
import { certificateRoot } from "../../lib/certificate";

// import { MdArrowDropDown } from "react-icons/md";
// import Dropdown from "../Dropdown";

export interface RoomMenuItem {
  key?: string;
  directoryListingId?: bigint;
  networkKey: Uint8Array;
  serverKey: Uint8Array;
  viewerCount: number;
  name: string;
}

export interface RoomMenuProps {
  onChange?: (item: ThreadProviderProps) => void;
}

enum Tab {
  Rooms,
  Whispers,
  Users,
}

interface TabsProps<T> {
  onChange: (tab: T) => void;
  active: T;
  tabs: { key: T; label: string }[];
}

const Tabs = <T extends number>({ onChange, active, tabs }: TabsProps<T>) => (
  <div className="room_meme__tab_bar">
    {tabs.map(({ key, label }) => (
      <button
        key={key}
        className={clsx("room_meme__tab", {
          "room_meme__tab--active": key === active,
        })}
        onClick={() => onChange(key)}
      >
        {label}
      </button>
    ))}
  </div>
);

export const RoomButtons: React.FC<RoomMenuProps> = ({ onChange }) => {
  const [activeTab, setActiveTab] = useState<Tab>(Tab.Rooms);
  const tabs = useMemo(
    () => [
      { key: Tab.Rooms, label: "rooms" },
      { key: Tab.Whispers, label: "whispers" },
      { key: Tab.Users, label: "users" },
    ],
    []
  );

  const list = (() => {
    switch (activeTab) {
      case Tab.Rooms:
        return <RoomsList onChange={onChange} />;
      case Tab.Whispers:
        return <WhispersList onChange={onChange} />;
      case Tab.Users:
        return <UsersList onChange={onChange} />;
    }
  })();

  return (
    <div className="room_meme">
      <Tabs onChange={setActiveTab} active={activeTab} tabs={tabs} />
      <div className="room_meme__content">{list}</div>
    </div>
  );
};

interface RoomsListItemProps extends RoomMenuProps {
  chat: RoomMenuItem;
}

const RoomsListItem: React.FC<RoomsListItemProps> = ({ onChange, chat }) => {
  const [, { openRoom }] = useChat();

  const handleClick = useStableCallback(() => {
    openRoom(chat.serverKey, chat.networkKey);
    onChange({ type: "ROOM", topicKey: chat.serverKey });
  });

  return (
    <button className="rooms_list__network_rooms_item" onClick={handleClick} key={chat.key}>
      <span className="rooms_list__network_rooms_item__name">{chat.name}</span>
      <span className="rooms_list__network_rooms_item__viewers">
        {chat.viewerCount.toLocaleString()}
        <HiOutlineUser />
      </span>
    </button>
  );
};

const RoomsList: React.FC<RoomMenuProps> = ({ onChange }) => {
  const networks = useContext(NetworkContext);
  const { directories } = useContext(DirectoryContext);
  const groups = useMemo(() => {
    const groups: { network: Network; rooms: RoomMenuItem[] }[] = [];
    for (const { network } of networks.items) {
      const networkKey = certificateRoot(network.certificate).key;
      const directory = directories[Base64.fromUint8Array(networkKey, true)];
      if (directory === undefined) {
        continue;
      }

      const rooms: RoomMenuItem[] = [];
      for (const { id, listing, viewerCount } of directory.listings.values()) {
        if (listing?.content?.case !== directoryv1.Listing.ContentCase.CHAT) {
          continue;
        }

        const { key, name } = listing.content.chat;
        rooms.push({
          key: Base64.fromUint8Array(key, true),
          directoryListingId: id,
          networkKey,
          serverKey: key,
          viewerCount,
          name,
        });
      }

      if (rooms.length === 0) {
        continue;
      }
      rooms.sort((a, b) => a.name.localeCompare(b.name));
      groups.push({ network, rooms });
    }

    return groups;
  }, [directories]);

  return (
    <Scrollbars autoHide={true} className="rooms_list">
      <div className="rooms_list__content">
        {groups.map(({ network, rooms }) => (
          <div key={network.id.toString()} className="rooms_list__network">
            <div className="rooms_list__network_name">
              {certificateRoot(network.certificate).subject}
            </div>
            <div className="rooms_list__network_rooms">
              {rooms.map((chat, i) => (
                <RoomsListItem key={i} chat={chat} onChange={onChange} />
              ))}
            </div>
          </div>
        ))}
      </div>
    </Scrollbars>
  );
};

const useMessageTimeFormatter = () => {
  // TODO: load formats from localization config
  return (time: Date) => {
    const now = new Date();
    const sameYear = now.getFullYear() === time.getFullYear();
    const sameMonth = sameYear && now.getMonth() === time.getMonth();
    const sameDay = sameMonth && now.getDate() === time.getDate();

    if (sameDay) {
      return date.format(time, "h:mm A");
    }
    if (sameYear) {
      return date.format(time, "MMM DD");
    }
    return date.format(time, "M/D/YY");
  };
};

interface WhispersListItemProps extends RoomMenuProps {
  thread: WhisperThread;
}

const WhispersListItem: React.FC<WhispersListItemProps> = ({ onChange, thread }) => {
  const [, { openWhispers }] = useChat();
  const formatMessageTime = useMessageTimeFormatter();

  // TODO: get these from somewhere meaningful... the thread? directory?
  const { items } = useContext(NetworkContext);
  const networkKeys = useMemo(
    () => items.map((i) => certificateRoot(i.network.certificate).key),
    [items]
  );

  const handleClick = useStableCallback(() => {
    openWhispers(thread.peerKey, networkKeys);
    onChange({ type: "WHISPER", topicKey: thread.peerKey });
  });

  return (
    <tr key={thread.id.toString()} className="whispers_list__row" onClick={handleClick}>
      <td className="whispers_list__label">
        <span className="whispers_list__alias">{thread.alias}</span>
        <span className="whispers_list__unread">{thread.unreadCount}</span>
      </td>
      <td className="whispers_list__time">
        {formatMessageTime(new Date(Number(thread.lastMessageTime)))}
      </td>
    </tr>
  );
};

const WhispersList: React.FC<RoomMenuProps> = ({ onChange }) => {
  const [{ whisperThreads }] = useChat();

  const sortedThreads = useMemo(
    () =>
      Array.from(whisperThreads.values()).sort((a, b) =>
        Number(b.lastMessageTime - a.lastMessageTime)
      ),
    [whisperThreads]
  );

  return (
    <Scrollbars autoHide={true} className="whispers_list">
      <table className="whispers_list__table">
        <tbody>
          {sortedThreads.map((thread, i) => (
            <WhispersListItem key={i} thread={thread} onChange={onChange} />
          ))}
        </tbody>
      </table>
    </Scrollbars>
  );
};

interface UsersListItemProps extends RoomMenuProps {
  networks: Uint8Array[];
  user: DirectoryUser;
}

const UsersListItem: React.FC<UsersListItemProps> = ({ onChange, networks, user }) => {
  const [, { openWhispers }] = useChat();

  const handleClick = useStableCallback(() => {
    openWhispers(user.peerKey, networks);
    onChange({ type: "WHISPER", topicKey: user.peerKey });
  });

  return (
    <button onClick={handleClick} key={user.id.toString()} className="user_list__item">
      {user.alias}
    </button>
  );
};

interface RoomUserThing extends RoomMenuProps {
  servers: directoryv1.Listing[];
  networks: Uint8Array[];
  user: DirectoryUser;
}

const UsersList: React.FC<RoomMenuProps> = ({ onChange }) => {
  const { directories } = useContext(DirectoryContext);
  const users = useMemo(() => {
    const users = new Map<string, RoomUserThing>();
    for (const { networkKey, listings } of Object.values(directories)) {
      for (const { listing, viewers } of listings.values()) {
        if (listing?.content?.case === directoryv1.Listing.ContentCase.CHAT) {
          for (const viewer of viewers.values()) {
            const key = Base64.fromUint8Array(viewer.peerKey, true);
            let user = users.get(key);
            if (user === undefined) {
              user = {
                servers: [],
                networks: [],
                user: viewer,
              };
              users.set(key, user);
            }
            user.servers.push(listing);
            user.networks.push(networkKey);
          }
        }
      }
    }
    return Array.from(users.values()).sort((a, b) => a.user.alias.localeCompare(b.user.alias));
  }, [directories]);

  return (
    <Scrollbars autoHide={true} className="user_list">
      <div className="user_list__content">
        {users.map(({ user, networks }, i) => (
          <UsersListItem key={i} user={user} networks={networks} onChange={onChange} />
        ))}
      </div>
    </Scrollbars>
  );
};

export const RoomList = forwardRef<HTMLDivElement, RoomMenuProps>((props, ref) => (
  <div ref={ref} className="room_list">
    <RoomButtons {...props} />
  </div>
));

// interface RoomDropdownPsop extends RoomMenuProps {
//   defaultSelection: RoomMenuItem;
// }

// export const RoomDropdown: React.FC<RoomDropdownPsop> = ({ onChange, defaultSelection }) => {
//   const [selection, setSelection] = useState<RoomMenuItem>(defaultSelection);
//   const handleChange = (item: RoomMenuItem) => {
//     setSelection(item);
//     onChange?.(item);
//   };

//   return (
//     <Dropdown
//       baseClassName="room_menu"
//       anchor={
//         <>
//           <div className="room_menu__text">{selection?.name}</div>
//           <MdArrowDropDown className="room_menu__icon" />
//         </>
//       }
//       items={<RoomButtons onChange={handleChange} />}
//     />
//   );
// };
