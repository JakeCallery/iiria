import React from 'react';
import ReactDOM from 'react-dom';
import WeatherUI from './components/weatherUI';

interface IAppProps {    
}


const App = (props: IAppProps): JSX.Element => {
    return (
        <div>
            <WeatherUI/>
        </div>
    )
}


ReactDOM.render(
    <App />, 
    document.querySelector('#root')
)