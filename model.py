import matplotlib.pyplot as plt
import matplotlib.image as mpimg
import numpy as np
import os

import sklearn
from numpy import argmax
from sklearn.metrics import ConfusionMatrixDisplay, confusion_matrix
from sklearn.model_selection import train_test_split
import tensorflow as tf
from tensorflow.python.keras.metrics import Metric
from tensorflow import keras
from tensorflow.keras import layers
from tensorflow.keras.models import Sequential
from tensorflow_core.python.keras.layers import Dense, Activation
from tensorflow_core.python.keras.optimizers import RMSprop
from tensorflow_core.python.keras.utils import np_utils

data = []
target = []


def readSingleImg(dir):
    img = mpimg.imread('ass2_processed_data/' + dir)
    img = img / 255
    return img


def readMultipleFile():
    files = os.listdir('ass2_processed_data')
    for f in files:
        img = readSingleImg(f)
        label = labelExtraction(f)
        rows, cols = img.shape
        img = img.reshape(rows * cols)
        data.append(img)
        target.append(label)


def labelExtraction(dir):
    label = dir[0]
    return onehotEncoding(label)


def onehotEncoding(char):
    y = 0
    if char == "A":
        y = 1.0

    if char == "B":
        y = 2.0

    if char == "C":
        y = 3.0

    if char == "D":
        y = 4.0

    if char == "E":
        y = 5.0

    if char == "F":
        y = 6.0

    if char == "G":
        y = 7.0

    if char == "H":
        y = 8.0

    if char == "J":
        y = 9.0

    if char == "K":
        y = 10.0

    if char == "L":
        y = 11.0

    if char == "M":
        y = 12.0

    if char == "N":
        y = 13.0

    if char == "P":
        y = 14.0

    if char == "Q":
        y = 15.0

    if char == "R":
        y = 16.0

    if char == "S":
        y = 17.0

    if char == "T":
        y = 18.0

    if char == "U":
        y = 19.0

    if char == "V":
        y = 20.0

    if char == "W":
        y = 21.0

    if char == "X":
        y = 22.0

    if char == "Y":
        y = 23.0

    if char == "Z":
        y = 24.0

    return y


def get_model():
    #  regression model, cannot use accuracy as the matrix
    model = Sequential([
        Dense(512, activation='relu', input_dim=1152),
        Dense(128, activation='relu', input_dim=1152),
        Dense(24, activation='softmax')
    ])
    # 定义优化器
    model.compile(
        optimizer='sgd',
        loss='mse',
        metrics=[tf.keras.metrics.MeanSquaredError()])
    return model


def ind2vec(ind, N=None):
    ind = np.asarray(ind)
    if N is None:
        N = ind.max() + 1
    return (np.arange(N) == ind[:, None]).astype(int)


def vec2ind(vec):
    y = []
    for v in vec:
        label = argmax(v)
        y.append(label)

    return np.array(y)


def train(x_train, y_train):
    model = get_model()
    model.fit(x_train, y_train, epochs=250)
    model.save('model/my_model')
    _, test_loss = model.evaluate(x_test, y_test)
    print(test_loss)


def load_model():
    model = keras.models.load_model("model/my_model")

    y_pred = model.predict(x_test)

    pred_class = np.argmax(y_pred, axis=1)
    y_class = np.argmax(y_test, axis=1)

    confusion_matrix = sklearn.metrics.confusion_matrix(y_true=y_class, y_pred=pred_class)  # shape=(12, 12)
    disp = ConfusionMatrixDisplay(confusion_matrix=confusion_matrix)
    disp.plot()
    plt.savefig('plot/python_cm.png')
    plt.show()
    print(confusion_matrix)

    # matrix = sklearn.metrics.confusion_matrix(y_test, np.rint(y_pred))
    # print(matrix)

    # test_loss, test_acc = model.evaluate(x_test, y_test)
    # print(test_loss)
    # print(test_acc)

    # label = np.array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23])
    # print(label)
    # cm = confusion_matrix(y_test, predictions, label)
    # disp = ConfusionMatrixDisplay(confusion_matrix=cm)
    # disp.plot()


if __name__ == '__main__':
    readMultipleFile()
    X = np.array(data)
    Y = np.array(target)
    x_train, x_test, y_train, y_test = train_test_split(X, Y, test_size=0.2)
    y_train = ind2vec(y_train, 24)
    y_test = ind2vec(y_test, 24)

    # print(y_test)
    # print(vec2ind(y_test))
    # print(vec2ind(y_test).shape)
    # print(argmax(y_test))
    # train(x_train, y_train)
    load_model()
