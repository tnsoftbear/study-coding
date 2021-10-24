import React from "react";
import { Square } from "../square/Square";

interface Props {
  squaresData: Array<string>,
  onClickButton: (i: number) => void
}

interface State {}

export default class Board extends React.Component<Props, State> {
  renderSquare(i: number) {
    return (
      <Square
        buttonNumber={i}
        buttonContent={this.props.squaresData[i]}
        onClickButton={() => this.props.onClickButton(i)}
      />
    );
  }

  render() {
    return (
      <div>
        <div className="board-row">
          {this.renderSquare(0)}
          {this.renderSquare(1)}
          {this.renderSquare(2)}
        </div>
        <div className="board-row">
          {this.renderSquare(3)}
          {this.renderSquare(4)}
          {this.renderSquare(5)}
        </div>
        <div className="board-row">
          {this.renderSquare(6)}
          {this.renderSquare(7)}
          {this.renderSquare(8)}
        </div>
      </div>
    );
  }
}
