import { useCallback } from "react";
import RequestResultWithFilmInfoFactory from "./RequestResultWithFilmInfoFactory";
import useRequest from "./UseRequest";
import { RequestResultInterface } from "./RequestResult";
import RequestResultWithFilmInfo from "./RequestResultWithFilmInfo";

const fetchFilmData = (idx: number) => {
  return fetch(`https://swapi.dev/api/films/${idx}/`)
    .then((response) => response.json())
    .then((data) => data);
};

export const useRequestResultWithFilmInfo = (id: number): RequestResultWithFilmInfo => {
  // const request = () => fetchFilmData(id + 1) - такое приводит к бесконечно повторяющимся запросам
  // потому что каждый раз вызывая useFilmInfo2() ф-ция request создаётся заного, 
  // она передаётся в качестве зависимости в useRequest
  // хук useEffect в useRequest видит, что зависимость [request] изменилась и перезапускает эффект.
  // 
  // useCallback запоминает значение функции которую передали и обновляет, когда id изменился
  const request = useCallback(() => fetchFilmData(id + 1), [id])
  const requestResult = useRequest(request) as RequestResultInterface;
  const requestResultWithFilmInfo = RequestResultWithFilmInfoFactory.createByRequestResult(requestResult);
  return requestResultWithFilmInfo;
};
