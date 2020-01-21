import React from "react";

let gaugePaths = [
    "M20.5377846,132.958454 L7.64718682,138.430675 C2.89926696,127.034525 0.199336394,114.571383 0.00703448198,101.500889 L14.008824,101.500347 C14.1992069,112.629508 16.5036318,123.243648 20.5377846,132.958454 Z ", // 1
    "M7.28254918,62.4565844 L20.1751083,67.928945 C16.2481288,77.6935876 14.0609701,88.345093 13.9974248,99.5 L0,99.5 L0.00939714293,98.3463172 C0.214928331,85.6697517 2.77923802,73.5665818 7.28254918,62.4565844 Z ", // 2
    "M28.0885261,30.5065948 L37.9892416,40.4082213 C30.8667873,47.8174822 25.0596284,56.5004342 20.9433251,66.0815172 L8.05309952,60.6092516 C12.8779107,49.3622999 19.7060288,39.1784383 28.0885261,30.5065948 Z ", // 3
    "M61.1721409,7.81598796 L66.4194863,20.8012858 C56.2624684,25.1127838 47.1053377,31.3211027 39.3925635,38.981773 L29.4930549,29.0821936 C38.5237178,20.1040302 49.2595535,12.8391547 61.1721409,7.81598796 Z ", //4
    "M98.4951114,0.011034482 L98.4956534,14.012824 C87.8346938,14.1951975 77.6463331,16.3174799 68.2683933,20.0418491 L63.0223981,7.05762986 C74.0207194,2.67855557 85.9794746,0.195172954 98.4951114,0.011034482 Z ", //5
    "M100.495,0.004 L99.996,0 C112.728957,0 124.90609,2.37976649 136.106888,6.71878852 L130.859784,19.704208 C121.42686,16.0760991 111.192899,14.0623609 100.496331,14.0014248 L100.495,0.004 Z", // 6
    "M170.016568,28.6058406 L160.115651,38.5050125 C152.287289,30.8506522 143.003217,24.6788917 132.717514,20.4438098 L137.963549,7.45960645 C150.005807,12.4055716 160.869561,19.6337312 170.016568,28.6058406 Z ", // 7
    "M191.790503,60.2651189 L178.899899,65.7373106 C174.701386,56.081869 168.782499,47.347271 161.530219,39.9204989 L171.430446,30.0205273 C179.943108,38.7090645 186.883301,48.9441026 191.790503,60.2651189 Z ", // 8
    "M199.994774,99.4994248 L185.991,99.5 L185.984478,98.5778327 C185.806909,87.625861 183.582093,77.1715545 179.676041,67.5809239 L192.567626,62.1085832 C197.296492,73.6494489 199.929938,86.2707922 199.994774,99.4994248 Z ", // 9
    "M199.983,101.5 L199.996,100 C199.996,113.750012 197.220879,126.85186 192.200176,138.776004 L179.309951,133.303726 C183.43341,123.495343 185.790535,112.7615 185.983176,101.500347 L199.983,101.5 Z ", // 10
];

let baseColor = "#dddddd";
let fullDelay = 0.25;

const colorFromCategory = {
    "high": "#f83e63",
    "medium": "#faba2f",
    "low": "#b2b2b2"
};

function getGaugeColor(score, category, index) {
    let fillColor = baseColor;
    let fillOpacity = 0.5;
    if (score >= index + 1) {
        fillColor = colorFromCategory[category];
        if (score !== 0) {
            let step = 1 / score;
            fillOpacity = step * (index + 1);
        }
    }
    return {fillColor, fillOpacity};
}

function RiskGauge({riskScore, category, shouldAnimate=true}) {

    let roundScore = Math.round(riskScore);
    let lowerCategory = category.toLowerCase();

    function renderText() {
        return <>
            <text fill={colorFromCategory[lowerCategory]} fontFamily="Poppins-SemiBold, Poppins" fontSize="90"
                  fontWeight="500">
                <tspan x={roundScore > 9 ? "60" : "73"} y="127">{roundScore}</tspan>
            </text>
            <text fill="#86888E" fontFamily="Poppins-Regular, Poppins" fontSize="16">
                <tspan x="61" y="184">risk score</tspan>
            </text>
            <text fill={colorFromCategory[lowerCategory]} fontFamily="HelveticaNeue-Medium, Helvetica Neue"
                  fontSize="27" fontWeight="400">
                <tspan x={lowerCategory === 'medium' ? "55" : lowerCategory === 'low' ? "78" : "74"}
                       y="163">{lowerCategory}</tspan>
            </text>
            <text fill="#BCBCBC" fontFamily="Poppins-Regular, Poppins" fontSize="10">
                <tspan x="84" y="196">{roundScore} of 10</tspan>
            </text>
        </>;
    }

    return (
        <div className={"RiskGauge"}>
            <svg xmlns="http://www.w3.org/2000/svg" width="210" height="203" viewBox="0 0 210 203">
                <g fill="none" fillRule="evenodd" transform="translate(5 3)">
                    {renderText()}
                    {gaugePaths.map((gaugePath, index) => {
                        let {fillColor, fillOpacity} = getGaugeColor(roundScore, lowerCategory, index);
                        let duration = (fullDelay / 10) * roundScore;
                        let delay = (roundScore !== 0) ? (duration / roundScore) * (index + 1) : 0;
                        return (
                            shouldAnimate ?
                                <path key={index} fill={baseColor} fillOpacity="0.5" fillRule="nonzero" d={gaugePath}>
                                    {roundScore >= index + 1 ?
                                        <>
                                            <animate attributeName="fill" from={fillColor} to={fillColor} begin={delay} dur="0.0001s" fill="freeze"/>
                                            <animate attributeName="fill-opacity" from="1" to={fillOpacity} begin={delay} dur={0.0001 + duration - delay} fill="freeze"/>
                                        </>
                                        :
                                        null
                                    }
                                </path>
                                :
                                <path key={index} fill={fillColor} fillOpacity={fillOpacity} fillRule="nonzero" d={gaugePath}/>
                        );
                    })}
                </g>
            </svg>
        </div>
    )
}

export default RiskGauge;