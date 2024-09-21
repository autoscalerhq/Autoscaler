"use client"
import axios from 'axios';

enum firstKnown {
    Youtube = "Youtube",
    Linkedin = "Linkedin",
    X = "X",
    Search_Engine = "Search Engine",
    Blogs = "Blogs",
    Review_Site = "Review Site",
    Word_Of_Mouth = "Word Of Mouth",
    Referral = "Referral",
    Other = "Other",
}

enum companySize {
    Only_Me = "Only Me",
    _2_5 = "2-5",
    _6_10 = "6-10",
    _11_50 = "11-50",
    _51_100 = "51-100",
    _100_plus = "100+",
}


const submit_hubspot_form = async (param: {email:string; firstname:string; lastname:string; firstLearn:firstKnown; companyName:string; companySize:companySize; phoneNumber:string;}) => {
    const portalId = '45292712';
    const formGuid = 'e3b417dd-1fe0-4d10-8be7-0fdd12f9da95';
    const config = { // important!
        headers: {
            'Content-Type': 'application/json',
        }
    }

    const response = await axios.post(`https://api.hsforms.com/submissions/v3/integration/submit/${portalId}/${formGuid}`,
        {
            portalId,
            formGuid,
            fields: [
                {
                    name: 'firstname',
                    value: param.firstname,
                },
                {
                    name: 'lastname',
                    value: param.lastname,
                },
                {
                    name: 'email',
                    value: param.email,
                },
                {
                    name: 'how_did_you_first_learn_about_us_',
                    value: param.firstLearn,
                },
                {
                    name: 'company',
                    value: param.companyName,
                },
                {
                    name: 'company_size',
                    value: param.companySize ,
                },
                {
                    name: 'mobilephone',
                    value: param.phoneNumber,
                },

            ],
        },
        config
    );
    return response;
}

export default function HomePage() {
    return (
        <>
            <div className={" flex items-center justify-center "}>
                <p>general</p>

            </div>
        </>
)
    ;
}

