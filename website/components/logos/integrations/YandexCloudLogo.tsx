import * as React from "react"
import {INTEGRATION_WIDTH} from "./constants";

const YandexCloudLogo = ({width = INTEGRATION_WIDTH}) => (
    <svg
        viewBox="0 0 120 120"
        xmlns="http://www.w3.org/2000/svg"
        xmlSpace="preserve"
        style={{
            fillRule: "evenodd",
            clipRule: "evenodd",
            strokeLinejoin: "round",
            strokeMiterlimit: 2,
        }}
        className="dark:text-white text-gray-900"
        width={width}
    >
        <path
            d="M109.4 60c0 27.3-22.1 49.4-49.4 49.4-7.144 0-12.479-2.682-15.147-5.934-2.668-3.253-3.419-8.575-2.5-16.263l15.9-3.154C72.78 81.388 80.784 73.306 83.65 58.817c.351-2.028.751-4.056 1.131-5.979.155-.787.307-1.556.45-2.3.433-2.352.851-4.587 1.235-6.636.404-2.16.769-4.114 1.073-5.783 1.627-9.419.305-16.374-3.354-21.212C99.239 25.368 109.4 41.489 109.4 60ZM60 10.6c-27.3 0-49.4 22.1-49.4 49.4 0 18.511 10.161 34.632 25.215 43.093-3.659-4.838-4.981-11.793-3.354-21.212.349-1.917.767-4.177 1.234-6.698l1.039-5.622c.143-.745.295-1.514.45-2.301.38-1.922.78-3.95 1.131-5.979 2.767-14.587 10.771-22.669 25.298-25.33l16.034-3.154c.919-7.688.168-13.01-2.5-16.263C72.479 13.282 67.144 10.6 60 10.6ZM120 60c0 33.1-26.9 60-60 60S0 93.1 0 60 26.9 0 60 0s60 26.9 60 60ZM75.745 43.836l-12.057 2.365c-10.277 1.972-15.021 6.703-16.997 16.953-.46 2.292-.86 4.465-1.247 6.565l-.334 1.813c-.297 1.577-.593 3.154-.89 4.632l11.958-2.365c10.277-1.971 15.12-6.702 17.097-16.953a421.01 421.01 0 0 0 1.234-6.505c.401-2.17.801-4.34 1.236-6.505Z"
            fill="currentColor"
        />
    </svg>
)

export default YandexCloudLogo
