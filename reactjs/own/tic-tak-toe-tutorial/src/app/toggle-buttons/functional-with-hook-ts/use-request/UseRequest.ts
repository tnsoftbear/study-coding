import { useState, useEffect, useMemo } from "react";
import { RequestResult } from "./RequestResult";

const useRequest = (request: () => Promise<RequestResult>): RequestResult => {
  // хук useMemo() кеширует результат выполнения функции, в отличии от useCallback(), которая кеширует значение - функцию
  const initialState = useMemo(() => RequestResult.constructorLoading(), []);
  const [dataState, setDataState] = useState<RequestResult>(initialState);
  useEffect(() => {
    setDataState(initialState);
    let cancelled = false;
    request()
      .then((data) => {
        if (!cancelled) {
          setDataState(new RequestResult(data));
        }
      })
      .catch((error) => {
        if (!cancelled) {
          setDataState(RequestResult.constructorError(error));
        }
      });
    return () => {
      cancelled = true;
    };
  }, [request, initialState]);
  return dataState;
};

export default useRequest;
