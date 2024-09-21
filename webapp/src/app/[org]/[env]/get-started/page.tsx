import React from 'react';

export default function getStarted() {
    return (
        <div
            className="flex w-full inset-0 h-full bg-white bg-[radial-gradient(#e5e7eb_1px,transparent_1px)] [background-size:16px_16px] "
        >
            <div className="space-y-4">
                <div className="space-y-4">
                    <div className="space-y-2">
                        {[
                            {
                                icon: (
                                    <svg
                                        stroke="currentColor"
                                        fill="currentColor"
                                        strokeWidth="0"
                                        viewBox="0 0 512 512"
                                        focusable="false"
                                        className="w-5 h-5"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M336 192h40a40 40 0 0140 40v192a40 40 0 01-40 40H136a40 40 0 01-40-40V232a40 40 0 0140-40h40m160-64l-80-80-80 80m80 193V48"
                                        ></path>
                                    </svg>
                                ),
                                text: 'Create a channel',
                            },
                            {
                                icon: (
                                    <svg
                                        stroke="currentColor"
                                        fill="currentColor"
                                        strokeWidth="0"
                                        viewBox="0 0 512 512"
                                        focusable="false"
                                        className="w-5 h-5"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M402 168c-2.93 40.67-33.1 72-66 72s-63.12-31.32-66-72c-3-42.31 26.37-72 66-72s69 30.46 66 72z"
                                        ></path>
                                        <path
                                            fill="none"
                                            strokeMiterlimit="10"
                                            strokeWidth="32"
                                            d="M336 304c-65.17 0-127.84 32.37-143.54 95.41-2.08 8.34 3.15 16.59 11.72 16.59h263.65c8.57 0 13.77-8.25 11.72-16.59C463.85 335.36 401.18 304 336 304z"
                                        ></path>
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M200 185.94c-2.34 32.48-26.72 58.06-53 58.06s-50.7-25.57-53-58.06C91.61 152.15 115.34 128 147 128s55.39 24.77 53 57.94z"
                                        ></path>
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeMiterlimit="10"
                                            strokeWidth="32"
                                            d="M206 306c-18.05-8.27-37.93-11.45-59-11.45-52 0-102.1 25.85-114.65 76.2-1.65 6.66 2.53 13.25 9.37 13.25H154"
                                        ></path>
                                    </svg>
                                ),
                                text: 'Identify users',
                            },
                            {
                                icon: (
                                    <svg
                                        stroke="currentColor"
                                        fill="currentColor"
                                        strokeWidth="0"
                                        viewBox="0 0 512 512"
                                        focusable="false"
                                        className="w-5 h-5"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M434.8 137.65l-149.36-68.1c-16.19-7.4-42.69-7.4-58.88 0L77.3 137.65c-17.6 8-17.6 21.09 0 29.09l148 67.5c16.89 7.7 44.69 7.7 61.58 0l148-67.5c17.52-8 17.52-21.1-.08-29.09zM160 308.52l-82.7 37.11c-17.6 8-17.6 21.1 0 29.1l148 67.5c16.89 7.69 44.69 7.69 61.58 0l148-67.5c17.6-8 17.6-21.1 0-29.1l-79.94-38.47"
                                        ></path>
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M160 204.48l-82.8 37.16c-17.6 8-17.6 21.1 0 29.1l148 67.49c16.89 7.7 44.69 7.7 61.58 0l148-67.49c17.7-8 17.7-21.1.1-29.1L352 204.48"
                                        ></path>
                                    </svg>
                                ),
                                text: 'Create a workflow',
                            },
                            {
                                icon: (
                                    <svg
                                        stroke="currentColor"
                                        fill="currentColor"
                                        strokeWidth="0"
                                        viewBox="0 0 512 512"
                                        focusable="false"
                                        className="w-5 h-5"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M53.12 199.94l400-151.39a8 8 0 0110.33 10.33l-151.39 400a8 8 0 01-15-.34l-67.4-166.09a16 16 0 00-10.11-10.11L53.46 215a8 8 0 01-.34-15.06zM460 52L227 285"
                                        ></path>
                                    </svg>
                                ),
                                text: 'Send a notification',
                            },
                            {
                                icon: (
                                    <svg
                                        stroke="currentColor"
                                        fill="currentColor"
                                        strokeWidth="0"
                                        viewBox="0 0 512 512"
                                        focusable="false"
                                        className="w-5 h-5"
                                        xmlns="http://www.w3.org/2000/svg"
                                    >
                                        <circle
                                            cx="129"
                                            cy="96"
                                            r="48"
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                        ></circle>
                                        <circle
                                            cx="129"
                                            cy="416"
                                            r="48"
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                        ></circle>
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M129 144v224"
                                        ></path>
                                        <circle
                                            cx="385"
                                            cy="288"
                                            r="48"
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                        ></circle>
                                        <path
                                            fill="none"
                                            strokeLinecap="round"
                                            strokeLinejoin="round"
                                            strokeWidth="32"
                                            d="M129 144c0 96 112 144 208 144"
                                        ></path>
                                    </svg>
                                ),
                                text: 'Push to production',
                            },
                        ].map((item, index) => (
                            <div key={index} className="flex items-center space-x-2">
                                <div className="p-2">{item.icon}</div>
                                <div>{item.text}</div>
                            </div>
                        ))}
                    </div>

                    <div className="space-y-2">
                        <div className="flex justify-between items-center">
                            <div className="text-lg font-bold">Create a channel</div>
                            <div className="bg-blue-500 text-white px-2 py-1 rounded">Dashboard</div>
                        </div>
                        <div className="text-sm">
                            <p>Create your first channel so you can start sending notifications with Knock.</p>
                        </div>
                        <div>
                            <a href="/bbd/integrations/channels" className="bg-blue-500 text-white px-4 py-2 rounded">
                                Create a channel
                            </a>
                        </div>
                        <div className="bg-gray-100 p-4 rounded space-y-2">
                            <div className="flex items-center space-x-2">
          <span role="img" aria-label="star">
            🌠
          </span>
                                <div>
                                    <p className="text-sm font-bold">Tip: get started with a preconfigured in-app
                                        channel</p>
                                    <p className="text-sm">
                                        Your account comes with an in-app feed channel pre-installed and ready to use.
                                        To
                                        get started
                                        building workflows that notify users in your product,{' '}
                                        <a href="/bbd/development/workflows" className="text-blue-500 underline">
                                            create your first workflow
                                        </a>
                                        .
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}