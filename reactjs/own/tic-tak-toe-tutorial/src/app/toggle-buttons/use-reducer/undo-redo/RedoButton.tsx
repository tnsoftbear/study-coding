import React from "react";

interface Props {
  dispatch: React.Dispatch<any>;
}

const RedoButton = ({dispatch}: Props) => {
  const redo = () => {
    dispatch({ type: 'redo' });
  };

  return (
    <button className="btn btn-info" onClick={redo}>
      Redo
    </button>
  );
};

export default RedoButton;
