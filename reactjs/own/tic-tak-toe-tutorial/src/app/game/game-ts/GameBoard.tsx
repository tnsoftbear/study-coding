import React from "react";
import Board from "./board/Board";
import SquareValue from "./square/SquareValue";
import Clock from "../../common/clock/Clock";
import WinnerDetector from "./domain/WinnerDetector";
// import { ll } from "./Debug.js";

interface History {
  squaresData: Array<string>;
}

interface Props {}
interface State {
  history: Array<History>;
  nextMoveChar: string;
  stepNumber: number;
}

export default class GameBoard extends React.Component<Props, State> {
 
  constructor(props: Props)
  {
    super(props);
    this.state = {
      history: [
        {
          squaresData: Array(9).fill(SquareValue.NONE),
        },
      ],
      nextMoveChar: SquareValue.IKS,
      stepNumber: 0,
    };
  }

  handleClick(i: number) {
    const state = this.state;
    const history = state.history.slice(0, state.stepNumber + 1);
    const current = history[history.length - 1];
    if (WinnerDetector.calculateWinner(current.squaresData)) {
      return;
    }
    if (!SquareValue.isNone(current.squaresData[i])) {
      return;
    }

    const squaresData = current.squaresData.slice(); // Clone array to be immutable
    squaresData[i] = state.nextMoveChar;
    const nextMoveChar = SquareValue.nextMove(state.nextMoveChar);
    const stepNumber = state.stepNumber + 1;
    const newState = {
      history: history.concat([
        {
          squaresData: squaresData,
        },
      ]),
      nextMoveChar,
      stepNumber,
    };
    this.setState(newState);
  }

  jumpTo(step: number) {
    this.setState({
      stepNumber: step,
      nextMoveChar: step % 2 === 0 ? SquareValue.IKS : SquareValue.ZERO,
    });
  }

  render() {
    const state = this.state;
    const history = state.history;
    const current = history[state.stepNumber];

    let gameStatus = null;
    const winner = WinnerDetector.calculateWinner(current.squaresData);
    if (winner) {
      gameStatus = "Winner is " + winner;
    } else {
      gameStatus = "Next player: " + state.nextMoveChar;
    }

    const moves = history.map((step, moveIndex) => {
      const desc = moveIndex ? "Go to move #" + moveIndex : "Go to game start";
      return (
        <li key={moveIndex}>
          <button onClick={() => this.jumpTo(moveIndex)}>{desc}</button>
        </li>
      );
    });

    return (
      <div className="game">
        <h1>Game by TS</h1>
        <div className="game-board">
          <Board
            squaresData={current.squaresData}
            onClickButton={(i) => this.handleClick(i)}
          />
        </div>
        <div className="game-info">
          <div className="status">{gameStatus}</div>
          <ol>{moves}</ol>
        </div>

        <div className="clock">
          <Clock textColor="red" />
        </div>
      </div>
    );
  }
}
