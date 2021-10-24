interface Props {
  buttonContent: string,
  buttonNumber: number,
  onClickButton: () => void
}

export function Square(props: Props) {
  const buttonContent = props.buttonContent;
  const id = "square" + props.buttonNumber;
  return (
    <button className="square" id={id} onClick={props.onClickButton}>
      {buttonContent}
    </button>
  );
}