import ToggleButton from "./ToggleButton";

interface Props {
  count: number;
}

const ToggleButtonList = ({ count }: Props) => {
  const buildListOfToggleButtons = () => {
    const toggleButtons = [];
    for (let i = 0; i < count; i++) {
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