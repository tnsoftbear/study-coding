import React from "react";

interface Props {
  dispatch: React.Dispatch<any>;
}

const AddToggleButton = ({dispatch}: Props) => {
  const addToggleButton = () => {
    dispatch({ type: 'add-toggle-button' });
  };

  return (
    <button className="btn btn-secondary" onClick={addToggleButton}>
      Add more
    </button>
  );
};

export default AddToggleButton;
