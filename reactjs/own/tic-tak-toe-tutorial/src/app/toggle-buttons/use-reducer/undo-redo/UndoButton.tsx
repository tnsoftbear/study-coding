import React from "react";

interface Props {
  dispatch: React.Dispatch<any>;
}

const UndoButton = ({dispatch}: Props) => {
  const undo = () => {
    dispatch({ type: 'undo' });
  };

  return (
    <button className="btn btn-success" onClick={undo}>
      Undo
    </button>
  );
};

export default UndoButton;
