import React, { useState } from "react";
import { AxiosError } from "axios";
import { fetchPageAnalysisData } from "../apiCalls";
import { ErrorResponseBody } from "../models/urlModels";

const Form = () => {
  const [url, setUrl] = useState("");

  const clickHandler = async (
    e: React.MouseEvent<HTMLFormElement, MouseEvent>
  ) => {
    e.preventDefault();
    try {
      const response = await fetchPageAnalysisData({ url });
      console.log(response);
    } catch (err) {
      const error = err as AxiosError<ErrorResponseBody>;
      console.log(error.response?.data);
    }
  };

  return (
    <form onClick={clickHandler}>
      <input
        type="text"
        name="url"
        placeholder="Enter url..."
        onChange={(e) => setUrl(e.target.value)}
      />
      <button type="submit">Analyze Page</button>
    </form>
  );
};

export default Form;
