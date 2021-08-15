import React, {useState} from 'react';
import { useEffect } from 'react';
import axios from 'axios';
import BigAnswer from './bigAnswer'
import WeatherDetails from './weatherDetails'

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
    humidity: number;
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
    humidity: -1.0,
} 

const WeatherUI = (props: IWeatherProps): JSX.Element => {

    const [weatherData, setWeatherData]:[IWeatherData, (weatherData: IWeatherData)=>void] = useState<IWeatherData>(defaultWeatherData);
    const [isLoading, setIsLoading]: [boolean, (isLoading: boolean) => void] = useState<boolean>(true);
    const [error, setError]: [string, (error: string) => void] = useState<string>(''); 

    //TODO: Set this up more properly based on running from npm start or npm build 
    let weatherAPIEndpoint: string = ""; 
    
    const apiServerHost = process.env.REACT_APP_API_SERVER_HOST;
    const apiServerPort = process.env.REACT_APP_API_SERVER_PORT;

    if(apiServerHost === "" && apiServerPort === "") {
        weatherAPIEndpoint = '/api/weather'
    } else {
        weatherAPIEndpoint = 'http://' + apiServerHost + ':' + apiServerPort + '/api/weather';
    }
    
    useEffect(() => {
        axios
            .get<IWeatherData>(
                weatherAPIEndpoint, 
                {
                    headers: {"Content-Type": "application/json"}
                })
            .then(response => {
                setWeatherData(response.data);
                setIsLoading(false);
            })
    }, [weatherAPIEndpoint]);

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
                    <WeatherDetails 
                        temp={weatherData.temperature.toString()}
                        weatherDesc={weatherData.weatherDesc}
                        uvIndex={weatherData.uvIndex.toString()}
                        uvIndexDesc={weatherData.uvDesc}
                        uvHealthDesc={weatherData.uvHealthDesc}
                        humidity={weatherData.humidity.toString()}
                    />
                </div>
            }
        </div>
    )

}

export default WeatherUI;