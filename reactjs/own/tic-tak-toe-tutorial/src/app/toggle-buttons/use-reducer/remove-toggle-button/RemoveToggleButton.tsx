import React from "react";

interface Props {
  length: number,
  dispatch: React.Dispatch<any>;
}

const RemoveToggleButton = ({length, dispatch}: Props) => {
  const removeToggleButton = () => {
    const removeIndex = Math.floor(Math.random() * length);
    dispatch({ type: 'remove-toggle-button', removeIndex });
  };

  return (
    <button className="btn btn-warning" onClick={removeToggleButton}>
      Remove random button
    </button>
  );
};

export default RemoveToggleButton;
