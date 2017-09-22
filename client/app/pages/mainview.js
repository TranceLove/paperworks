"use strict"

import React from "react";
import { Text, View, Card } from 'react-native';
import { Button } from 'react-native-elements'
import { Styles } from './../styles/styles'
import MapView from 'react-native-maps';

export default ({navigation}) => (
    <MapView style={Styles.MainView.map}/>
)
