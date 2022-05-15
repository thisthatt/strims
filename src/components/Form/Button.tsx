// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import "./Button.scss";

import clsx from "clsx";
import React, { ComponentProps } from "react";

export const ButtonSet: React.FC = ({ children }: ButtonProps) => (
  <div className="input_label input_buttonset">{children}</div>
);

export type ButtonProps = ComponentProps<"button">;

const Button: React.FC<ButtonProps> = ({ children, className, ...inputProps }: ButtonProps) => (
  <button className={clsx("input", "input_button", className)} {...inputProps}>
    {children}
  </button>
);

export default Button;
