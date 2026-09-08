<?php

namespace App\View\Components;

use Closure;
use Illuminate\Contracts\View\View;
use Illuminate\View\Component;

class BadgeSks extends Component
{
    public function __construct(public int $sks)
    {
    }

    public function render(): View|Closure
    {
        return view('components.badge-sks');
    }
}