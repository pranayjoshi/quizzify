import React, { useState, useEffect } from "react";
import Navbar from "../Layout/navbar";
import { Link } from "react-router-dom";
import { useParams, useNavigate } from "react-router-dom";
import { useLocation } from 'react-router-dom';
import { useSelector } from "react-redux";

const QuizPage = () => {
  const quiz_id = useParams();
  console.log(quiz_id)

  const name = useSelector((state) => state.name);
  const token = useSelector((state) => state.token);
  const [marks, setMarks] = useState(0);
  const [answers, setAnswers] = useState({});
  const navigate = useNavigate();
  
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
  
    // Prepare the results data
    const results = {
      name: name, // Replace with the actual user's name
      quiz_name: token, // Replace with the actual quiz name
      quiz_id: quiz_id.quiz_id, // Replace with the actual quiz ID
      scored_marks: correctAnswers,
      total_marks: totalMarks,
      Author: name // Replace with the actual quiz author
    };
    console.log(results);
  
    // Send the results data to the server
    fetch(`http://localhost:8082/post_results?quiz_id=${quiz_id.quiz_id}`, { // Replace with the actual server URL
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(results)
    })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
    navigate("/");
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