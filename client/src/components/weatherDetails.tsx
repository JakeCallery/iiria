import React from 'react';

interface IDetailsProps {
    temp: string;
    weatherDesc: string;
    uvIndex: string;
    uvIndexDesc: string;
    uvHealthDesc: string;
    humidity: string;
}

const WeatherDetails = ({temp, weatherDesc, uvIndex, uvIndexDesc, uvHealthDesc, humidity}: IDetailsProps): JSX.Element => {
    return (
        <div>
            <ul>
                <li>Temp: {temp}</li>
                <li>Weather: {weatherDesc}</li>
                <li>Humidity: {humidity}%</li>
                <li>UVIndex: {uvIndex} ({uvIndexDesc})</li>
                <li>UV Concern: {uvHealthDesc}</li>
            </ul>
            

        </div>
    )
}

export default WeatherDetails;