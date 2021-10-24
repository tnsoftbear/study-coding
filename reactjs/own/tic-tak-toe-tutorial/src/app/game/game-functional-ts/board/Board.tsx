import { Square } from "../square/Square";

interface Props {
  squaresData: Array<string>;
  onClickButton: (i: number) => void;
}

const Board = (props: Props) => {
  const renderSquare = (i: number) => {
    return (
      <Square
        buttonNumber={i}
        buttonContent={props.squaresData[i]}
        onClickButton={() => props.onClickButton(i)}
      />
    );
  };

  return (
    <div>
      <div className="board-row">
        {renderSquare(0)}
        {renderSquare(1)}
        {renderSquare(2)}
      </div>
      <div className="board-row">
        {renderSquare(3)}
        {renderSquare(4)}
        {renderSquare(5)}
      </div>
      <div className="board-row">
        {renderSquare(6)}
        {renderSquare(7)}
        {renderSquare(8)}
      </div>
    </div>
  );
};

export default Board;
