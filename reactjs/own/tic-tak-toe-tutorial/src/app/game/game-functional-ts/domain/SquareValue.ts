export default class SquareValue {
  static NONE = "_";
  static IKS = "X";
  static ZERO = "O";

  static isNone(value: string) {
    return value === SquareValue.NONE;
  }

  static nextMove(value: string) {
    return value === SquareValue.IKS ? SquareValue.ZERO : SquareValue.IKS;
  }
}
