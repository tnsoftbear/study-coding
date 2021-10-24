import { ButtonStateInterface } from "../ButtonState";

interface Props {
  buttonStates: Array<ButtonStateInterface>;
  setButtonStates: (buttonStates: Array<ButtonStateInterface>) => void;
}

const DropRandomButton = ({buttonStates, setButtonStates}: Props) => {
  const dropRandomToggleButton = () => {
    const droppedIndex = Math.floor(Math.random() * buttonStates.length);
    buttonStates[droppedIndex].toggledCount = 0;
    buttonStates[droppedIndex].isToggle = false;
    setButtonStates([...buttonStates]);
  };

  return (
    <button className="btn btn-danger" onClick={dropRandomToggleButton}>
      Drop random
    </button>
  );
};

export default DropRandomButton;
