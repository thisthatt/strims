import "./Dropdown.scss";

import clsx from "clsx";
import React, { useRef } from "react";
import { useToggle } from "react-use";

import useClickAway from "../../hooks/useClickAway";

export interface DropdownProps {
  baseClassName?: string;
  anchor: React.ReactNode;
  items: React.ReactNode | React.ReactNode[];
}

const Dropdown: React.FC<DropdownProps> = ({ baseClassName, anchor, items }) => {
  const ref = useRef<HTMLDivElement>(null);

  const [open, toggleOpen] = useToggle(false);
  useClickAway(ref, () => toggleOpen(false), ["click"]);

  return (
    <div className={clsx("dropdown", baseClassName)} ref={ref} onClick={() => toggleOpen()}>
      <div className={clsx("dropdown__anchor", `${baseClassName}__anchor`)}>{anchor}</div>
      <div
        className={clsx({
          "dropdown__menu": true,
          "dropdown__menu--open": open,
          [`${baseClassName}__menu`]: true,
          [`${baseClassName}__menu--open`]: open,
        })}
      >
        {items}
      </div>
    </div>
  );
};

export default Dropdown;
