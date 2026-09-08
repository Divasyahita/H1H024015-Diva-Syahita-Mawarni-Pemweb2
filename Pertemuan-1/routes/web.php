<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MatakuliahController;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/data-matakuliah', [MatakuliahController::class, 'index'])
    ->name('matakuliah.index');

Route::get('/data-matakuliah/{kode}', [MatakuliahController::class, 'show'])
    ->name('matakuliah.show');