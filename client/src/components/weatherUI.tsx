import React, {useState} from 'react';
import { useEffect } from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';
import BigAnswer from './bigAnswer'
import WeatherDeatils from './weatherDetails'

interface IWeatherProps {

}

interface IWeatherData {
    temperature: number;
    preciptationType: number;
    weatherCode: number;
    uvIndex: number;
    uvHealthConcern: number;
    weatherDesc: string;
    uvDesc: string;
    uvHealthDesc: string;
    precipitationDesc: string;
}

const defaultWeatherData: IWeatherData = {
    temperature: -1,
    preciptationType: -1,
    weatherCode: -1,
    uvIndex: -1,
    uvHealthConcern: -1,
    weatherDesc: "none",
    uvDesc: "none",
    uvHealthDesc: "none",
    precipitationDesc: "none",
} 

const WeatherUI = (props: IWeatherProps): JSX.Element => {

    const [weatherData, setWeatherData]:[IWeatherData, (weatherData: IWeatherData)=>void] = useState<IWeatherData>(defaultWeatherData);
    const [isLoading, setIsLoading]: [boolean, (isLoading: boolean) => void] = useState<boolean>(true);
    const [error, setError]: [string, (error: string) => void] = useState<string>(''); 

    useEffect(() => {
        axios
            .get<IWeatherData>(
                'http://localhost:9090/weather', 
                {
                    headers: {"Content-Type": "application/json"}
                })
            .then(response => {
                setWeatherData(response.data);
                setIsLoading(false);
            })
    }, []);

    const isRaining = (data: IWeatherData): boolean => {
        console.log("data.weatherCode: ", data.weatherCode)
        if(data.weatherCode >= 4000 && data.weatherCode <= 4999) {
            return true;
        } else {
            return false;
        }
    }

    return (
        <div>
            {isLoading && <h1>Loading Data...</h1>}
            {!isLoading && 
                <div>
                    <h1>Is It Raining in Austin?</h1>
                    <BigAnswer isRaining={isRaining(weatherData)}/>
                    <WeatherDeatils 
                        temp={weatherData.temperature.toString()}
                        weatherDesc={weatherData.weatherDesc}
                        uvIndex={weatherData.uvIndex.toString()}
                        uvIndexDesc={weatherData.uvDesc}
                        uvHealthDesc={weatherData.uvHealthDesc}
                    />
                </div>
            }
        </div>
    )

}

export default WeatherUI;