import React, {useState} from 'react';
import { useEffect } from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';

interface AppProps {
    
}

interface IWeatherData {
    temperature: number;
    preciptationType: number;
    weatherCode: number;
    uvIndex: number;
    uvHealthConcern: number;
}

const defaultWeatherData: IWeatherData = {
    temperature: -1,
    preciptationType: -1,
    weatherCode: -1,
    uvIndex: -1,
    uvHealthConcern: -1,
} 



const App = (props: AppProps): JSX.Element => {
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

    return (
        <div>
            {isLoading && <div>Loading...</div>}
            {weatherData.temperature}
        </div>
    )
}


ReactDOM.render(
    <App />, 
    document.querySelector('#root')
)