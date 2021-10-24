export default class SquareValue {
  static NONE = "_";
  static IKS = "X";
  static ZERO = "O";

  static isNone(value) {
    return value === SquareValue.NONE;
  }

  static nextMove(value) {
    return value === SquareValue.IKS ? SquareValue.ZERO : SquareValue.IKS;
  }
}
