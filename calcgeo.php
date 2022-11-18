<?php
    $raio=5;
    $largura=2;
    $altura=4;
    $pi = 3.14;
    $areaCirc= $pi * pow($raio,2);
    $circunf=2*$pi*$raio;
    $areaRect=$largura*$altura;
    $perim=2*($largura+$altura);
    echo "Area do Circulo = $areaCirc <br><br>";
    echo "Circunferencia do Circulo = $circunf <br><br>";
    echo "Area do Retangulo = $areaRect <br><br>";
    echo "Perimetro do Retangulo = $perim";
?>
