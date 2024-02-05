import React from "react";
import { useParams, Link, useNavigate } from "react-router-dom";

const ShowQuizIdPage = () => {
  const { quiz_id } = useParams(); // Get the quiz_id parameter from the URL
  const navigate = useNavigate();

  return (
    <div className="min-h-screen flex items-center justify-center ">
      <div className="p-6 bg-gray-800 flex items-center justify-center px-8 py-16 border border-gray-200 rounded-lg">
        <h1 className="text-2xl font-bold text-white">Quiz ID:</h1>
        <pre className="text-xl text-white mx-4 bg-gray-900 p-3 rounded-md">
          <code>{quiz_id}</code>
        </pre>
        <button
          onClick={() => {navigator.clipboard.writeText(quiz_id!); navigate("/")}}
          className=" px-4 py-2 bg-white border lg:text-lg rounded-full text-white bg-gradient-to-r from-green-400 to-blue-500 focus:outline-none"
        >
          Copy
        </button>
      </div>
    </div>
  );
};

export default ShowQuizIdPage;
