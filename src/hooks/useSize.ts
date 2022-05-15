// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import useResizeObserver from "@react-hook/resize-observer";
import * as React from "react";

const useSize = (target: HTMLElement): DOMRectReadOnly => {
  const [size, setSize] = React.useState<DOMRectReadOnly>();

  React.useLayoutEffect(() => {
    setSize(target?.getBoundingClientRect());
  }, [target]);

  useResizeObserver(target, (entry) => setSize(entry.contentRect));
  return size;
};

export default useSize;
