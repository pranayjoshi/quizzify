import React, { useState, useEffect } from "react";
import Navbar from "../Layout/navbar";
import { Link } from "react-router-dom";
import { useParams } from "react-router-dom";
import { useLocation } from 'react-router-dom';

// const dummyData = Array.from({ length: 10 }, (_, i) => ({
//   question: `Question ${i + 1}`,
//   options: ["Option 1", "Option 2", "Option 3", "Option 4"],
//   answer: "Option 1",
// }));

const QuizPage = () => {
  const [marks, setMarks] = useState(0);
  const [answers, setAnswers] = useState({});
  
  const location = useLocation();
  var questions = location.state.content;
  
  questions = questions.Quiz
  const totalMarks = questions.length;

  const handleAnswerChange = (questionIndex, selectedOption) => {
    setAnswers(prevAnswers => ({
      ...prevAnswers,
      [questionIndex]: selectedOption
    }));
  };

  const handleSubmit = () => {
    let correctAnswers = 0;

    questions.forEach((question, index) => {
      if (answers[index] === question.Answer) {
        correctAnswers++;
      }
    });

    console.log(`Total correct answers: ${correctAnswers}`);
  };
  return (
    <div className="h-screen overflow-y-auto justify-center  bg-gray-900 text-white">
      <Navbar />
      <div className="w-1/2 mx-auto">
        <div>
          <div className="p-6  text-white">
            <h1 className="text-3xl font-bold">Quick Quiz</h1>
          </div>
          {questions.map((question, index) => (
          <div key={index} className="p-6 m-3 rounded-2xl shadow-md bg-gray-800 ">
            <h2 className="text-2xl font-bold mb-2">{question.Question}</h2>
            {[question.Option1, question.Option2, question.Option3, question.Option4].map((option, i) => (
        <div key={i} className="mb-2">
          <label className="inline-flex items-center">
            <input
              type="radio"
              className="form-radio text-green-500"
              name={`question-${index}`}
              value={option}
              onChange={() => handleAnswerChange(index, option)}
            />
            <span className="ml-2">{option}</span>
          </label>
        </div>
      ))}
          </div>
        ))}
        </div>
        <div className="flex justify-between p-3">
          <Link
            to="/"
            className="text-blue-500 hover:text-blue-700 bg-white font-bold py-2 px-4 rounded-full border border-blue-400"
          >
            Back
          </Link>
          <button className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2" onClick={handleSubmit}>
            Submit
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuizPage;
