import React, { useState, useRef } from "react";
import { AxiosError } from "axios";
import ValidResponse from "./ValidResponse";
import InvalidResponse from "./InvalidResponse";
import { fetchPageAnalysisData } from "../apiCalls";
import { ErrorResponseBody, ResponseBody } from "../models/urlModels";

const Form = () => {
  const inputRef = useRef<HTMLInputElement>(null);
  const [responseData, setResponseData] = useState<
    ResponseBody | ErrorResponseBody | undefined
  >(undefined);

  const clickHandler = async (
    e: React.MouseEvent<HTMLFormElement, MouseEvent>
  ) => {
    e.preventDefault();
    try {
      const response = await fetchPageAnalysisData({
        url: inputRef.current?.value || "",
      });
      setResponseData(response.data);
    } catch (err) {
      const error = err as AxiosError<ErrorResponseBody>;
      setResponseData(error.response?.data);
    } finally {
      inputRef.current!.value = "";
      inputRef.current?.focus();
    }
  };

  let urlDataComponent;
  if (responseData) {
    if (!(responseData as ErrorResponseBody).err_message) {
      urlDataComponent = (
        <ValidResponse urlResponse={responseData as ResponseBody} />
      );
    } else {
      urlDataComponent = (
        <InvalidResponse errResponse={responseData as ErrorResponseBody} />
      );
    }
  }

  return (
    <div className="content">
      <form onSubmit={clickHandler}>
        <input
          type="text"
          name="url"
          placeholder="Enter url..."
          ref={inputRef}
        />
        <button type="submit">Analyze Page</button>
      </form>
      {urlDataComponent}
    </div>
  );
};

export default Form;
