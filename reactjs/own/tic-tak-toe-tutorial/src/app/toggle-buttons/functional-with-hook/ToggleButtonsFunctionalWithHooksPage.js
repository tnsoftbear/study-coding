import React, { useState, useEffect } from "react";
import ToggleButtonBuilder from "./ToggleButtonBuilder.js";
import ButtonState from "./ButtonState.js";
//import { ll } from "../../common/debug/Debug.js";

export const ToggleButtonContext = React.createContext();

const ToggleButtonsFunctionalWithHooksPage = (props) => {

  const TITLE = "Toggle buttons with useState() hook"; 

  // Для инициализации используется колбек, чтобы инициализация не вызывалась повторно при каждом перерендеринге компонента.
  const [pageState, setPageState] = useState(
    () => ({ 
      mainTitle: TITLE,
      subTitle: '',
      data: Array.from({ length: props.count }, () => new ButtonState()),
    })
  );

  // Обновить через колбек ф-цию на основе предыдущего состояния
  const toggle = (buttonIndex) => (newIsToggle) => {
    setPageState((prev) => applyToggle(prev, buttonIndex, newIsToggle));
  };

  // Колбек для применения изменений на предыдущем состоянии
  const applyToggle = (prevButtonStates, buttonIndex, newIsToggle) => {
    const { data } = prevButtonStates;
    data.map((buttonState) => {
      buttonState.isToggle = false;
      return buttonState;
    });
    data[buttonIndex].isToggle = newIsToggle;
    data[buttonIndex].toggledCount += newIsToggle;
    const sum = sumToggleOn(data);
    return {
      ...prevButtonStates,
      data,
      subTitle: `Sum is ${sum}`
    };
  };

  const sumToggleOn = (data) => {
    return data.reduce((accumulator, current) => accumulator + current.toggledCount, 0);
  }

  useEffect(() => {
    console.log(`Title changed to ${pageState.subTitle}`)
  }, [pageState.subTitle]);

  const buildListOfToggleButtons = () => {
    const toggleButtons = [];
    for (let i = 0; i < props.count; i++) {
      const ToggleButton = ToggleButtonBuilder();
      toggleButtons[i] = <ToggleButton idx={i} doToggle={toggle(i)} />;
    }
    const listToggleButtons = toggleButtons.map((tb) => (
      <li key={tb.props.idx}>{tb}</li>
    ));
    return listToggleButtons;
  };

  return (
    <div>
      <h1>{pageState.mainTitle}</h1>
      <h2>{pageState.subTitle}</h2>
      <ToggleButtonContext.Provider value={pageState.data}>
        <ol>{buildListOfToggleButtons()}</ol>
      </ToggleButtonContext.Provider>
      <pre>{JSON.stringify(pageState, null, 2)}</pre>
    </div>
  );
};

export default ToggleButtonsFunctionalWithHooksPage;

// Простой вариант обновления состояния без использования колбек-функции
// 
// const toggleSimple = (buttonIndex) => (newIsToggle) => {
//   buttonStates.map((buttonState) => {
//     buttonState.isToggle = false;
//     return buttonState;
//   });
//   buttonStates[buttonIndex].isToggle = newIsToggle;
//   buttonStates[buttonIndex].toggledCount += newIsToggle;
//   setButtonStates([...buttonStates]);
// };
