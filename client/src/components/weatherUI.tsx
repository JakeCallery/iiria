import React, {useState} from 'react';
import { useEffect } from 'react';
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

    const apiServerHost = process.env.REACT_APP_API_SERVER_HOST;
    const apiServerPort = process.env.REACT_APP_API_SERVER_PORT;
    //const weatherAPIEndpoint = 'http://' + apiServerHost + ':' + apiServerPort + '/api/weather';
    const weatherAPIEndpoint = '/api/weather';
    console.log(weatherAPIEndpoint)
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