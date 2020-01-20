import React from "react";
import Modal from "../Modal/Modal";
import "./RiskModal.scss"
import RiskGauge from "./RiskGauge";
import cn from "classnames";

let riskTypes = [
    "Basic",
    "Aggravation",
    "Remediation",
];


function RiskModal({riskData, closePopup}) {
    function renderRisksDescription() {
        const descShorts = [];
        riskData.risk.riskItems.forEach(r => {
            if (r.type !== "Remediation"){
                r.shortDescription.split('\n').forEach(desc => descShorts.push(desc));
            }

        });
        let unique = [...new Set(descShorts)];
        let shortUniqueList = Array.from(unique.values()).slice(0, 4);

        return <div className='risks-description'>
            <div className='risks-warning-img'/>
            <div className='risks-texts'>
                <span>Risks:</span>
                <div className='texts-list'>
                    {shortUniqueList.map((riskIDesc, i) => {
                        return (
                            <div key={i}>- {riskIDesc}</div>
                        );
                    })}
                </div>
            </div>
        </div>;
    }

    function renderRisksDetails() {
        riskData.risk.riskItems.sort((item1, item2) => {
            const typeDiff = riskTypes.indexOf(item1.type) - riskTypes.indexOf(item2.type);
            if (typeDiff !== 0)
                return typeDiff;
            return (item1.score - item2.score) * (item1.type === 'Remediation' ? 1 : -1);
        });
        return (<>
            { riskData.risk.riskItems.map((riskItem, i) => {
                let score = riskItem.score.toString();
                let barItemsToColor = riskItem.score;
                if (riskItem.type !== 'Basic') {
                    const percent = (riskItem.score * 100 - 100);
                    barItemsToColor = Math.abs(percent) / 10;
                    const sign = percent >= 0 ? "+" : "";
                    score = sign + percent + "%";
                }
                let shortDesc = riskItem.title || riskItem.shortDescription || riskItem.description.substring(0, 60) + "...";

                const barIndexes = [];
                for (let barIndex = 1; barIndex <= 10; barIndex++) {
                    barIndexes.push(barIndex);
                }

                return (
                    <div key={i} className='riskItem'>
                        <div className='riskItemScore'>
                            <span className={`${riskItem.type.toLowerCase()}`}>{score}</span>
                        </div>
                        <div className='riskItemBar'>
                            <div className='fullBar'>
                                {
                                    barIndexes.map((val, barIndex) => {
                                        const classesArray = ['barItem'];
                                        if (val <= barItemsToColor) {
                                            classesArray.push(`${riskItem.type.toLowerCase()}`)
                                        } else if (val - barItemsToColor < 1) {
                                            classesArray.push(`${riskItem.type.toLowerCase()}`);
                                            classesArray.push('halfBar');
                                        }
                                        return (
                                            <div key={barIndex} className={cn(classesArray)}/>
                                        );
                                    })
                                }
                            </div>
                            <div className={"short-desc"}>{shortDesc}<div className={"tooltip-hint"}><div className={"tooltiptext"}>{riskItem.description}</div></div></div>
                        </div>
                    </div>
                );
            })}
        </>);
    }

    return (
        <Modal onBackdropClick={closePopup} modalClassNames="RiskModal">
            <div className='title-row'>
                <div className='name'>{riskData.kind} / {riskData.name}</div>
                <div className='close-wrapper'>
                    <div className='close oc-icon close-icon' onClick={closePopup}></div>
                </div>
            </div>
            <div className='main-panel'>
                <div className='main-panel-top'>
                    <div className='gauge-panel'>
                        <RiskGauge riskScore={riskData.risk.riskScore} category={riskData.risk.riskCategory}/>
                    </div>
                    <div className={'workloadRiskItems'}>
                        {renderRisksDetails()}
                    </div>
                </div>
                <div className='main-panel-bottom'>
                    {renderRisksDescription()}
                </div>
            </div>
        </Modal>
    );
}

export default RiskModal;