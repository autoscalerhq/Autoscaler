import React from 'react';

const YourPage = () => {
    return (
        <div style={{ padding: '24px', background: 'rgb(41, 41, 51)', borderRadius: '8px', color: 'rgb(130, 130, 153)' }}>
            <div className="css-f1ix14 e1h3gpue0">
                <div>
                    <div className="mantine-Group-root mantine-tx9rmw">
                        <div className="mantine-Group-root mantine-1dib7ec">
                            {/* The SVG goes here */}
                            <h1 className="mantine-Text-root mantine-Title-root mantine-ydgrn4">
                                Role-based access control
                            </h1>
                        </div>
                    </div>
                    <div className="mantine-Text-root mantine-1brs9lt">
                        Securely manage users' permissions to access system resources.
                    </div>
                </div>
                <div className="mantine-Group-root mantine-1ifsn9r">
                    <button className="mantine-UnstyledButton-root mantine-Button-root mantine-1g21dij" type="button" data-button="true">
                        <div className="mantine-3xbgk5 mantine-Button-inner">
                            <span className="mantine-Button-label mantine-ec3lso">
                                <div className="mantine-Group-root mantine-1dib7ec">
                                    {/* The SVG goes here */}
                                    Schedule a call
                                </div>
                            </span>
                        </div>
                    </button>
                    <button className="mantine-UnstyledButton-root mantine-ActionIcon-root mantine-c15loc" type="button">
                        {/* The SVG goes here */}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default YourPage;