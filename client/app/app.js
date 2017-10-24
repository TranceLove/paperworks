import React from 'react';
import { Routes, RoutesDev } from './routes';

export default class PaperWorkApp extends React.Component
{
    render()
    {
        const isDev = true
        return (isDev ? (<RoutesDev/>) : (<Routes/>))
    }
}
