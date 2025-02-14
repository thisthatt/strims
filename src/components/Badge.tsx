// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import "./Badge.scss";

import React from "react";

export interface BadgeProps {
  count: number;
  max?: number;
}

const Badge: React.FC<BadgeProps> = ({ count, max = count }) => (
  <span className="badge">
    {Math.min(count, max).toLocaleString()}
    {count > max && "+"}
  </span>
);

export default Badge;
