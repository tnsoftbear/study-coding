import { useContext } from "react";
import { ToggleButtonContext } from "../ToggleButtonContext";
import ToggleButton from "./ToggleButton";

const ToggleButtonList = () => {
  
  const { appState } = useContext(ToggleButtonContext);

  const buildListOfToggleButtons = () => {
    const toggleButtons = [];
    for (let i = 0; i < appState.buttonStates.length; i++) {
      toggleButtons[i] = <ToggleButton idx={i} />;
    }
    const listToggleButtons = toggleButtons.map((tb) => (
      <li key={tb.props.idx}>{tb}</li>
    ));
    return listToggleButtons;
  };

  return (
    <>
      <ol>{buildListOfToggleButtons()}</ol>
    </>
  );
};

export default ToggleButtonList;