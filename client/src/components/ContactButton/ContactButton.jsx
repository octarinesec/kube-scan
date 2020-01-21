import React from 'react';

import './ContactButton.scss';

function ContactButton ({contactLink}) {
    return (
        <span className="contact-button">
            <a href={contactLink} target="_blank" rel="noopener noreferrer">Contact us</a>
        </span>
    );
}

export default ContactButton;


