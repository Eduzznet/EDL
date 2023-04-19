<?php
function quick_sort($vet)
{
    if(count($vet) <= 1){
        return $vet;
    }
    else{
        $aux = $vet[0];
        $esquerda = array();
        $direita = array();
        for($i = 1; $i < count($vet); $i++)
        {
            if($vet[$i] < $aux){
                $esquerda[] = $vet[$i];
            }
            else{
                $direita[] = $vet[$i];
            }
        }
        return array_merge(quick_sort($esquerda), array($aux), quick_sort($direita));
    }
}
$desordenado = array(5,3,8,6,2,7,-1,9,11,87,-18,22,13,0);
echo 'Vetor original: ';
echo implode(",",$desordenado)."<br>";
$ordenado = quick_sort($desordenado);
echo 'Vetor ordenado: ';
echo implode(",",$ordenado)."<br>";
?>