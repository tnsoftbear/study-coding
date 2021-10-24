import { useState, useEffect } from "react";
import { useFilmInfo } from "../film-info/UseFilmInfoHook";
import { useRequestResultWithFilmInfo } from "../use-request/UseRequestResultWithFilmInfo";

interface Props {
  clickedButtonIndex: number | null;
}

const ClickedButtonInfo = (props: Props) => {
  const [message, setMessage] = useState<string>("Not clicked yet");

  useEffect(() => {
    if (props.clickedButtonIndex !== null) {
      const someRandom = Math.round(
        Math.random() * Math.pow(10, props.clickedButtonIndex + 1)
      );
      const message = `The last clicked button index is "${props.clickedButtonIndex}" and random is "${someRandom}"`;
      setMessage(message);
    }
  }, [props.clickedButtonIndex]);

  const filmInfoRequestResult = useFilmInfo(props.clickedButtonIndex);

  const id = props.clickedButtonIndex === null ? 0 : props.clickedButtonIndex;
  const requestResultWithFilmInfo = useRequestResultWithFilmInfo(id);

  return (
    <div>
      <p> {message} </p>
      <p>{filmInfoRequestResult.render()}</p>
      <p>{requestResultWithFilmInfo.render()}</p>
    </div>
  );
};

export default ClickedButtonInfo;
