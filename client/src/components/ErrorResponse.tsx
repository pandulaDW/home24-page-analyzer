import React from "react";
import { ErrorResponseBody } from "../models/urlModels";

interface Props {
  errResponse: ErrorResponseBody;
}

const ValidResponse: React.FC<Props> = ({ errResponse }) => {
  return (
    <div className="invalid-response">
      <h2>
        {errResponse.err_status_code === 400
          ? "Invalid Request !!"
          : "Server Error !!"}
      </h2>
      <p>
        <span>{errResponse.err_message} - </span>
        <span>{errResponse.err_description}</span>
      </p>
    </div>
  );
};

export default ValidResponse;
