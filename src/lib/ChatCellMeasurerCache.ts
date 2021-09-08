import { CellMeasurerCacheParams } from "react-virtualized";
import { CellMeasurerCacheInterface } from "react-virtualized/dist/es/CellMeasurer";

// ChatCellMeasurerCache minimally implements the react-virtualized interface
// with a prune function to sync the cache indices with the message history.
export default class ChatCellMeasurerCache implements CellMeasurerCacheInterface {
  readonly defaultHeight = 0;
  readonly defaultWidth = 0;

  private values: number[] = [];

  columnWidth: (params: { index: number }) => number = () => 0;
  rowHeight: (params: { index: number }) => number;

  constructor(params?: CellMeasurerCacheParams) {
    this.rowHeight = (params) => this.values[params.index] || 0;
  }

  prune(n: number): void {
    this.values.splice(0, n);
  }

  clear(rowIndex: number, columnIndex: number): void {
    this.values[rowIndex] = undefined;
  }

  clearAll(): void {
    this.values = [];
  }

  hasFixedHeight(): boolean {
    return false;
  }

  hasFixedWidth(): boolean {
    return true;
  }

  getHeight(rowIndex: number, columnIndex: number): number {
    return this.values[rowIndex];
  }

  getWidth(rowIndex: number, columnIndex: number): number {
    return 0;
  }

  has(rowIndex: number, columnIndex: number): boolean {
    return this.values[rowIndex] !== undefined;
  }

  set(rowIndex: number, columnIndex: number, width: number, height: number): void {
    this.values[rowIndex] = height;
  }
}
