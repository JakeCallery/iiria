import React from 'react';
import ReactDOM from 'react-dom';

interface IDetailsProps {
    temp: string;
    weatherDesc: string;
    uvIndex: string;
    uvIndexDesc: string;
    uvHealthDesc: string;
}

const WeatherDetails = ({temp, weatherDesc, uvIndex, uvIndexDesc, uvHealthDesc}: IDetailsProps): JSX.Element => {
    return (
        <div>
            <ul>
                <li>Temp: {temp}</li>
                <li>Weather: {weatherDesc}</li>
                <li>UVIndex: {uvIndex} ({uvIndexDesc})</li>
                <li>UV Concern: {uvHealthDesc}</li>
            </ul>
            

        </div>
    )
}

export default WeatherDetails;