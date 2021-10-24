import { useState, useEffect } from "react";

interface Props {
  textColor: string
}

const ClockFunctional = (props: Props) => {
  const divStyle = {
    color: props.textColor ?? "blue",
  };
  const [date, setDate] = useState(new Date().toLocaleTimeString());

  useEffect(() => {
    const timerId = setInterval(() => {
      setDate(new Date().toLocaleTimeString());
    }, 1000);
    return function cleanUp() {
      clearInterval(timerId);
    };
  }, []);


  return (
    <div>
      <h1 style={divStyle}>{date}</h1>
    </div>
  );
};

export default ClockFunctional;