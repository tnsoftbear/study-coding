import React from "react";

interface Props {
  length: number;
  dispatch: React.Dispatch<any>;
}

const DropRandomButton = ({length, dispatch}: Props) => {
  const dropRandomToggleButton = () => {
    const droppedIndex = Math.floor(Math.random() * length);
    dispatch({ type: 'drop-random', droppedIndex });
  };

  return (
    <button className="btn btn-danger" onClick={dropRandomToggleButton}>
      Unset random state
    </button>
  );
};

export default DropRandomButton;
