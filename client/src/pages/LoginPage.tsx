import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { setToken } from "../state";
import Navbar from "../Layout/navbar";
import { useNavigate } from 'react-router-dom';


export const LoginPage = () => {
  const [name, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const dispatch = useDispatch();
  const navigate = useNavigate();
  
  const handleSubmit = async (event: any) => {
    event.preventDefault();
    
    const user = { name, password };
    console.log(user);
    try {
      const response = await fetch('http://localhost:8082/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
      });
      var data = await response.json();
        dispatch(setToken({ token: data.token, name: data.name}));
      console.log(data);

      if (response.status === 200) {
        console.log(`Successfully logged in user: ${name}`);
      } else {
        console.error(`Error: ${response.status}`);
      }
    } catch (error) {
      console.error(error);
    }
    navigate("/");
  };

  return (
    <div className="h-screen">
      <Navbar/>
      <div className="flex items-center justify-center h-full">
        <div className="flex items-center justify-center px-4 py-10 sm:px-6 lg:px-8 sm:py-16 lg:py-24 border border-gray-300 rounded-md bg-gray-800 m-auto text-white">
        <div className="xl:w-full xl:max-w-lg 2xl:max-w-xl xl:mx-auto tex ">
            <h2 className="text-3xl font-bold leading-tight text-white sm:text-4xl">
              Sign In
            </h2>

            <form id="my-form" className="mt-8" onSubmit={handleSubmit}>
              <div className="space-y-5 w-full">
                <div>
                  <label
                    htmlFor="username"
                    className="text-base font-medium text-gray-200"
                  >
                    {" "}
                    Username{" "}
                  </label>
                  <div className="mt-2.5">
                    <input
                      name="name"
                      value={name}
                      onChange={(e) => setUsername(e.target.value)}
                      className="flex h-10 w-full rounded-md border border-gray-300 bg-transparent py-2 px-3 text-sm placeholder:text-gray-400 focus:outline-none focus:ring-1 focus:ring-gray-400 focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-700 dark:text-gray-50 dark:focus:ring-gray-400 dark:focus:ring-offset-gray-900"
                      placeholder="Enter Your Username"
                    ></input>
                  </div>
                </div>
                <div>
                  <label
                    htmlFor="password"
                    className="text-base font-medium text-gray-200"
                  >
                    {" "}
                    Password{" "}
                  </label>
                  <div className="mt-2.5">
                    <input
                      name="password"
                      type="password"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      className="flex h-10 w-full rounded-md border border-gray-300 bg-transparent py-2 px-3 text-sm placeholder:text-gray-400 focus:outline-none focus:ring-1 focus:ring-gray-400 focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-50 dark:border-gray-700 dark:text-gray-50 dark:focus:ring-gray-400 dark:focus:ring-offset-gray-900"
                      placeholder="Enter Your Password"
                    ></input>
                  </div>
                </div>
                <div>
                <p>
                  <span className="text-gray-500 dark:text-gray-400 text-sm">
                    Read our{" "}
                    <span className="capitalize text-indigo-600">
                      privacy policy
                    </span>{" "}
                    and{" "}
                    <span className="capitalize text-indigo-600">
                      terms of service
                    </span>{" "}
                    to learn more
                  </span>
                </p>
              </div>

                <div className="flex justify-between mt-8">
                  <button className="px-8 py-2  border-white border-2 text-white rounded-full">
                    back
                  </button>
                  <button
                    type="submit"
                    id="my-form"
                    className="text-white bg-gradient-to-r from-green-400 to-blue-500 rounded-full px-6 py-2"
                  >
                    Submit
                  </button>
                </div>
                
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;