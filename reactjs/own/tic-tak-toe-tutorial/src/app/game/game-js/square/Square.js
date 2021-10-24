import React from "react";

export function Square(props) {
  const buttonContent = props.buttonContent;
  const id = "square" + props.buttonNumber;
  return (
    <button className="square" id={id} onClick={props.onClickButton}>
      {buttonContent}
    </button>
  );
}