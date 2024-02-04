import React from 'react';
import { useParams, Link } from 'react-router-dom';

const ShowQuizIdPage = () => {
  const { quiz_id } = useParams(); // Get the quiz_id parameter from the URL

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-800">
      <h1 className="text-xl text-white mr-4">Quiz ID: {quiz_id}</h1>
      <Link to="/" className="mt-4 px-4 py-2 bg-white border  lg:text-lg rounded-full text-white bg-gradient-to-r from-green-400 to-blue-500 focus:outline-none">Back</Link>
    </div>
  );
};

export default ShowQuizIdPage;