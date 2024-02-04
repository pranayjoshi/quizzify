

import React from "react";
import { Link } from "react-router-dom";
import Navbar from "../Layout/navbar";

const CodeEnterPage = () => {

  return (
    <div className="h-screen bg-inherit items-center justify-center">
      <Navbar />
      <div className="flex items-center justify-center justify-items-center h-full w-full">
  <div className="bg-gray-800  rounded-md p-8 items-center  w-7/12">
    <h1 className="text-2xl font-bold mr-4 text-white">Enter Code</h1>
    <div className="flex mt-4 w-full">
    <input name="workspace" id="workspace" type="text" className="rounded-full px-4 py-2 border border-gray-300 focus:outline-none bg-gray-700 text-white w-full lg:text-xl " placeholder="Enter quiz code..."/>
    <a target="_self" id="wsb" className="rounded-full ml-4 px-3 py-3 lg:px-8 lg:py-6 text-white bg-gradient-to-r from-green-400 to-blue-500 focus:outline-none">
    <p className="text-xl rounded-full"> Start!</p>
    </a>
    </div>
    <div className="lg:h-8 h-4"></div>
    <hr/>
    <button  className="mt-4 px-4 py-2 bg-white border  lg:text-lg rounded-full text-white bg-gradient-to-r from-green-400 to-blue-500 focus:outline-none">Back</button>
  </div>
</div>
    </div>
  );
};

export default CodeEnterPage;