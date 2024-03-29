import React from 'react';

interface IBigAnswerProps {
    isRaining: boolean;
}

const BigAnswer = ({isRaining}: IBigAnswerProps): JSX.Element => {
    
    const getAnswer = (isRaining:Boolean):string => {
        if(isRaining) {
            return "YES!"
        } else {
            return "NO!"
        }
    }

    return (
        <h1>{getAnswer(isRaining)}</h1>
    )
}

export default BigAnswer;