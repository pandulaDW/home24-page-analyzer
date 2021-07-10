import React from "react";
import { ErrorResponseBody } from "../models/urlModels";

interface Props {
  errResponse: ErrorResponseBody;
}

const ValidResponse: React.FC<Props> = ({ errResponse }) => {
  return (
    <div>
      <h1>
        {errResponse.err_status_code === 400
          ? "Invalid Request"
          : "Server Error"}
      </h1>
      <ul>
        <li>
          <span>Error Title</span>
          <span>{errResponse.err_message}</span>
        </li>
        <li>
          <span>Error Description</span>
          <span>{errResponse.err_description}</span>
        </li>
      </ul>
    </div>
  );
};

export default ValidResponse;
