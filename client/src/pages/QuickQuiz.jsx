import React, { useState, useEffect } from "react";
import Navbar from "../Layout/navbar";
import { Link } from "react-router-dom";

const dummyData = [
  {
    question: "What is the capital of France?",
    options: ["Paris", "London", "Berlin", "Madrid"],
    answer: "Paris",
  },
  {
    question: "Who wrote 'To Kill a Mockingbird'?",
    options: ["Harper Lee", "George Orwell", "Mark Twain", "Ernest Hemingway"],
    answer: "Harper Lee",
  },
  {
    question: "What is the square root of 144?",
    options: ["10", "12", "14", "16"],
    answer: "12",
  },
  {
    question: "What is the chemical symbol for Hydrogen?",
    options: ["H", "He", "Hy", "Ho"],
    answer: "H",
  },
  {
    question: "Who painted the Mona Lisa?",
    options: ["Vincent van Gogh", "Pablo Picasso", "Leonardo da Vinci", "Claude Monet"],
    answer: "Leonardo da Vinci",
  },
  {
    question: "What is the largest planet in our solar system?",
    options: ["Earth", "Mars", "Jupiter", "Saturn"],
    answer: "Jupiter",
  },
  {
    question: "What is the capital of Australia?",
    options: ["Sydney", "Melbourne", "Canberra", "Perth"],
    answer: "Canberra",
  },
  {
    question: "Who is the author of '1984'?",
    options: ["George Orwell", "Aldous Huxley", "Ray Bradbury", "Isaac Asimov"],
    answer: "George Orwell",
  },
  {
    question: "What is the square root of 256?",
    options: ["14", "16", "18", "20"],
    answer: "16",
  },
  {
    question: "What is the chemical symbol for Oxygen?",
    options: ["O", "Ox", "Og", "Om"],
    answer: "O",
  },
];
const QuickQuiz = () => {
  const [questions, setQuestions] = useState([]);

  useEffect(() => {
    setQuestions(dummyData);
  }, []);

  return (
    <div className="h-screen overflow-y-auto justify-center  bg-gray-900 text-white">
      <Navbar />
      <div className="w-1/2 mx-auto">
        <div>
          <div className="p-6  text-white">
            <h1 className="text-3xl font-bold">Quick Quiz</h1>
          </div>
          {questions.map((question, index) => (
            <div
              key={index}
              className="p-6 m-3 rounded-2xl shadow-md bg-gray-800 "
            >
              <h2 className="text-2xl font-bold mb-2">{question.question}</h2>
              {question.options.map((option, i) => (
                <div key={i} className="mb-2">
                  <label className="inline-flex items-center">
                    <input
                      type="radio"
                      className="form-radio text-green-500"
                      name={`question-${index}`}
                      value={option}
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
          <button className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2">
            Submit
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuickQuiz;
